# Website Searches


## July 30, 2024

Tools we want to look into:
- https://github.com/go-chi/chi


Resources I used:
- https://www.reddit.com/r/golang/comments/16d54ae/how_i_can_call_method_of_one_package_under/
- https://medium.com/@maciek.pilot2/golang-project-structure-b88327220d73 


Creating portable packages in golang
- https://go.dev/doc/tutorial/create-module
- https://stackoverflow.com/questions/52026284/accessing-local-packages-within-a-go-module-go-1-11 

- https://ukiahsmith.com/blog/a-gentle-introduction-to-golang-modules/




How to reach local go packages in another files [x](https://go.dev/doc/tutorial/call-module-code)
1) make sure its up on github 
2) run `go mod init github.com/endingwithali/fitnessapp/backend`
3) use the following command to let go know to look at local files
`go mod edit --replace github.com/endingwithali/fitnessapp/backend=../backend`



Using https://github.com/joho/godotenv to get env variables 
- How to learn about OidC connection and the handshake that occurs:
https://connect2id.com/learn/openid-connect


better explanation of oauth and passing around sso 
https://www.bacancytechnology.com/blog/sso-using-oauth-in-golang-application




https://www.mongodb.com/blog/post/quick-start-golang--mongodb--modeling-documents-with-go-data-structures
https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents
https://www.mongodb.com/blog/post/quick-start-golang--mongodb--modeling-documents-with-go-data-structures

https://github.com/markbates/goth

https://github.com/markbates/goth/blob/dcd837795bb225182b7ca6a6ac86ca4e79ace5ec/examples/main.go#L11
https://blog.seriesci.com/how-to-mock-oauth-in-go/

https://medium.com/bigcommerce-developer-blog/how-to-test-app-authentication-locally-with-ngrok-149150bfe4cf

https://discord.com/developers/docs/topics/oauth2

https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-fiber-version-4la0



10/17/2024
- https://github.com/karankumarshreds/GoAuthentication
Example of login auth session store using gorilla in golang 




## 10/22/2024
I AM CONFUSED with the following - how to safely store that a user has been logged in and how to use session cookies using gorilla to do so

What I know / where im missing info 
THE LOGIN:
1) user initiates auth using discord + oauth2 proceedure by starting the call via /auth/login using goth library to kick off auth handling 
2) discord sends back relevant info to /auth/login/callback and auth is completed using goth library
    - in the goth library call to completeuserauth -> call to StoreInSession method is called
    - StoreInSession method 
    ```go
        const SessionName untyped string = "_gothic_session"
        func StoreInSession(key string, value string, req *http.Request, res http.ResponseWriter) error {
            session, _ := Store.New(req, SessionName)

            if err := updateSessionValue(session, key, value); err != nil {
                return err
            }

            return session.Save(req, res)
        }
    ```

Given what Goth does in the method completeuserauth, so I as an implementer need to write code that does session management - if so how - or can i use a call like completeuserauth to validate a user is already logged in?


T.TV/USUALLYHIGH: you store a unique string in the cookie, can be anything, a uuid4. In your session store, for example Redis, you store this unique token with the corresponding user id. On future request, you would retrieve the unique token from the cookie, get associated user id, then retrieve the user details from your database


Coldfloat: session_token -> user_id -> call db -> user details

HexxCon: In php, I do the following. 1. Store session when user successfully logs in. 2. Assign identifier to session, such as userid. 3. Create a function that checks if session value has been assigned, if so user is logged in, if not redirect to login page. 4. Destroy session on logout and clear the session var

sudomateo: I was lurking a bit but yeah pretty much what HexxCon said. Store session with something to uniquely identify the user and read it on middleware to determine if used is logged in.


```sql
user_DB: 
    database: _id for the user gen by db
              discord_id - from discord oauth value to validate active discord user thing 
```        

i dont want to store the _id in the session ! that is unique for the database and should stay only for interacting between database values 

uidsecurity: (this is ACCURATE)
external user_id -> -> _id -> usertable_id in other tables  

### THIS IS THE ANSWER
UsuallyHigh: You can store sessions in your database too, doesn't have to be Redis. Don't overthink it. Create a new "sessions" table, which has 2 columns. "session_token" and "user_id". Those are your active sessions. There's more to it, such as expiration, but worry about it later. On successful login, you generate a new "session_token" for the user, and persist it. Now you can put the "session_token" in a cookie, ideally with HttpOnly and Secure flags. You then know if the user is logged on future request


so to follow up:
once user is logged in