// This event is fired each time the user updates the text in the omnibox,
// as long as the extension's keyword mode is still active.
chrome.omnibox.onInputChanged.addListener(
  function(text, suggest) {
    console.log('inputChanged: ' + text);
    suggest([
      {content: text + " one", description: "the first one"},
      {content: text + " number two", description: "the second entry"}
    ]);
  });

const AWS = require('aws-sdk.js');
const endpoint = new AWS.Endpoint("https://localhost:8000");
const config = require('config.js');
AWS.config.update(config.aws_local_config);
const dynamodb = new AWS.DynamoDB.DocumentClient();

// This event is fired with the user accepts the input in the omnibox.
chrome.omnibox.onInputEntered.addListener(
  function(text) {
    console.log('inputEntered: ' + text);
    alert('You just typed "' + text + '"');
    var params = {
        TableName: config.aws_table_name,
        Key: {
            'KEY_NAME': { Name: text }
        },
        ProjectionExpression: 'ATTRIBUTE_NAME'
    }

  /*  dynamodb.batchGetItem(params, function (err, data) {
      if (err) console.log(err, err.stack); // an error occurred
      else     chrome.tabs.update({url: 'https://www.apple.com'})           // successful response
    });*/
    chrome.tabs.update({url: 'https://www.apple.com'})
  });
