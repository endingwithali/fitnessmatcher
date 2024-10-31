# Backend for the fitness app

## Requirements


## Running 

How to run
``` 
go run cmd/app.go
```
Once running, in a new terminal window, run
```
ngrok http 3000
```
And update oauth2 redirect uri with uri created 

https://discord.com/developers/applications

to the uri created, with 
[redirecturi]/auth/login/callback

and update the .env value for `redirectURL`

when running locally, instead of going to localhost, go to the ngrok url


Now you can build and install that program with the go tool:
$ go install cmd/app