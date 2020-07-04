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
    newurl = 'http://localhost:8080/load?short=' + text;
    fetch(newurl, {mode: 'cors'})
        .then(function(response) {
          if (!response.ok) {
                console.log('there is a problem. Status Code: ' + response.status);
                return;
          };
          response.json().then(function(data) {
            console.log('link: ' + data);
            chrome.tabs.update({url: data})
          });
        })
        .catch(function(error) {
          console.log('Looks like there was a problem: \n', error);
        });
  });
