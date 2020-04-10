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


// This event is fired with the user accepts the input in the omnibox.
chrome.omnibox.onInputEntered.addListener(
  function(text) {
    alert('You just typed "' + text + '"');
    newurl = 'http://localhost:8080/load?short=' + text;
    alert('newurl is ' + newurl)
    fetch(newurl, {mode: 'cors'})
        .then(function(response) {
            console.log(response.headers.get('Content-Type'));
            console.log(response.headers.get('Date'));
            console.log(response.status);
            console.log(response.statusText);
            console.log(response.url);
          if (!response.ok) {
                console.log('there is a problem. Status Code: ' + response.status);
                alert('problem! ' + response.status);
                return;
          };
          response.json().then(function(data) {
            alert ('actually normal');
            console.log('link: ' + data);
            chrome.tabs.update({url: data})
          });
        })
        .catch(function(error) {
          console.log('Looks like there was a problem: \n', error);
        });
    //chrome.tabs.update({url: 'https://www.example.com'})
  });
