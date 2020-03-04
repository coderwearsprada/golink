/*chrome.omnibox.onInputEntered.addListener(
    function(text) {
        alert("testing");
        chrome.tabs.getSelected(null, function(tab)
        {
            var url;
            if (text.substr(0, 3) == 'go/') {
                url = 'https://bing.com/q/' + text;
            // If text does not look like a URL, user probably selected the default suggestion, eg reddit.com for your example
            } else {
                url = text;
            }
            url = 'https://bing.com/q/' + text;
            navigate(url)
        });
    }
 );*/
alert("here!")
console.log("where console is")

chrome.omnibox.onInputChanged.addListener(
  function(text, suggest) {
    suggest([
      {content: text + "/top/?sort=top&t=all", description: text + "/top/ (all time)"},//all time top posts
      {content: text + "/controversial/?sort=top&t=all", description: text + "/controversial/ (all time)"},// controversial
      {content: text + "/new/", description: text + "/new/"}// new posts
    ]);
  });

chrome.omnibox.onInputEntered.addListener(
  function(text) {

    if (text.indexOf("/") < 1) {
      text += "/";
    }
    if (text.indexOf("http") < 0) {
      text = "http://our-internal-portal/" + text;
    }
    alert('We are taking you to: "' + text + '"');
    navigate(text);
});

function navigate(url) {
  chrome.tabs.getSelected(null, function(tab) {
    chrome.tabs.update(tab.id, {url: url});
  });
}