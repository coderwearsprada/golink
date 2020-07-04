# golink

This project provides you a way to create your customized short link for any URL you want to remember.

a very preliminary draft:

Launch the server.go with
go run server.go

your server will run on port 8080 on localhost

then you load the extension to your Chrome browser

prerequisite - you have a dynamodb local server running on your localhost

Step 1:
you go to localhost:8080, and add a bunch of names and urls you want to create map for

e.g. short name: git, target url: https://github.com/coderwearsprada

Step 2:
go to your browser, type "goto" followed by a space, then type the short name you setup already.
e.g. "goto" + space + "git"

Now you are on your target url


