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