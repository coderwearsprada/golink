package main

import (
    //"bytes"
    "fmt"
    //"io"
    "encoding/json"
    "net/http"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    //"github.com/gorilla/mux"
    //"io/ioutil"
)

func headers (w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

type Item struct {
    Name   string
    Link   string
}

type response struct {
    status  string
    link   string
}

func load (w http.ResponseWriter, req *http.Request) {
    fmt.Println("GET params were:", req.URL.Query())
    short := req.URL.Query().Get("short")
    if short == "" {
        fmt.Println("need a short name!")
        return
    } else {
        fmt.Println("short name is " + short)
    }
    sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
    svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")})
    tableName := "GoMap"

    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Name": {
                S: aws.String(short),
            },
        },
    })
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    item := Item{}
    err = dynamodbattribute.UnmarshalMap(result.Item, &item)

    if err != nil {
        panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
        return
    }

    if item.Name == "" {
        fmt.Println("Could not find " + short)
        return
    }
    fmt.Println("Successfully found " + item.Name + " pointing to " + item.Link)

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(item.Link)
}

func createTable (w http.ResponseWriter, req *http.Request) {
    sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
    svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")})
    tableName := "GoMap"
    input := &dynamodb.CreateTableInput{
        AttributeDefinitions: []*dynamodb.AttributeDefinition{
            {
                AttributeName: aws.String("Name"),
                AttributeType: aws.String("S"),
            },
        },
        KeySchema: []*dynamodb.KeySchemaElement{
            {
                AttributeName: aws.String("Name"),
                KeyType:       aws.String("HASH"),
            },
        },
        ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
            ReadCapacityUnits:  aws.Int64(10),
            WriteCapacityUnits: aws.Int64(10),
        },
        TableName: aws.String(tableName),
    }

    _, err := svc.CreateTable(input)
    if err != nil {
        fmt.Println("Got error calling CreateTable:")
        fmt.Println(err.Error())
    }

    fmt.Println("Created the table", tableName)
}

func updateLink (w http.ResponseWriter, req *http.Request) {
    short := req.URL.Query().Get("short")
    if short == "" {
        fmt.Println("need a short name!")
        return
    } else {
        fmt.Println("short name is " + short)
    }
    link := req.URL.Query().Get("link")
    if link == "" {
        fmt.Println("need a link!")
        return
    } else {
        fmt.Println("link is " + link)
    }

    sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-2")})
    svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")})
    tableName := "GoMap"
    fmt.Println("Updating " + short + " to point to " + link)
    input := &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                S: aws.String(link),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Name": {
                S: aws.String(short),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Link = :r"),
    }

    _, err := svc.UpdateItem(input)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("Successfully updated " + short + " to point to " + link)
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
}

func main() {
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/load", load)
    http.HandleFunc("/update", updateLink)
    http.HandleFunc("/createTable", createTable)

    http.ListenAndServe(":8080", nil)
}