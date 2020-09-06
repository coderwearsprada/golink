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
    "log"
//    "golang.org/x/crypto/acme/autocert"
//   "golang.org/x/net/http2"
    "strings"
)

func headers (w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

type Item struct {
    Short  string
    Link   string
    Owner  string
}

type response struct {
    status  string
    link   string
}

var sess = session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
// Create DynamoDB client
var svc = dynamodb.New(sess)

var tableName = "go_link_dev"

func load (w http.ResponseWriter, req *http.Request) {
    fmt.Println("GET params were:", req.URL.Query())
    short := req.URL.Query().Get("short")
    if short == "" {
        fmt.Println("need a short name!")
        return
    } else {
        fmt.Println("short name is " + short)
    }

    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Short": {
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

    if item.Short == "" {
        fmt.Println("Could not find " + short)
        return
    }
    fmt.Println("Successfully found " + item.Short + " pointing to " + item.Link)

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(item.Link)
}

func getMine (w http.ResponseWriter, req *http.Request) {
    fmt.Println("GET params were:", req.URL.Query())

    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Owner": {
                S: aws.String("zhanyun"),
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

    if item.Short == "" {
        fmt.Println("Could not find any")
        return
    }
    fmt.Println("Successfully found " + item.Short + " pointing to " + item.Link)

    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(item.Link)
}

func createTable (w http.ResponseWriter, req *http.Request) {
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

    req.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    fmt.Println(req.Form) // print information on server side.
    fmt.Println("path", req.URL.Path)
    fmt.Println("scheme", req.URL.Scheme)
    fmt.Println(req.Form["url_long"])
    short := ""
    link := ""
    for k, v := range req.Form {
        if k == "short" {
            short = strings.Join(v, "")
        }
        if k == "link" {
            link = strings.Join(v, "")
        }
    }

    // Log
    fmt.Println("short:", short)
    fmt.Println("link:", link)
    fmt.Println("Updating " + short + " to point to " + link)

    input := &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                S: aws.String(link),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Short": {
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
    http.HandleFunc("/get-mine", getMine)
    http.Handle("/", http.FileServer(http.Dir("./assets/")))
    err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
    if err != nil {
        log.Fatal("ListenAndServeTLS: ", err)
    }
}