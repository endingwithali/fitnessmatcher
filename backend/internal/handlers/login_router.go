package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/mongo"

	// "golang.org/x/oauth2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
)

/*
Login Flow pulled from here
https://github.com/ravener/discord-oauth2/blob/master/example/main.go
https://github.com/golang/oauth2/blob/master/google/example_test.go

Things to understand next:
- How to use oauth2 login for account creation on our end
  - how does oauth authentication lead to user creation on a backend server

- how, after login, do we persist authentication and authorization once the user is created?
  - jwt information storing happens ? maybe?
*/
// var state = "EVENTUALLY MAKE A RANDOM GENERATOR"

type LoginRouterConfig struct {
	database mongo.Database
	ctx      context.Context
	// oauth2 oauth2.Config
}

func gothSetup() {
	key := os.Getenv("SESSION_SECRET") // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30               // 30 days
	isProd := false

	redirectURL := os.Getenv("OAUTH_REDIRECT") // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.GetProviderName = func(req *http.Request) (string, error) {
		return "discord", nil
	}
	gothic.Store = store
	goth.UseProviders(discord.New(os.Getenv("DISCORD_KEY"), os.Getenv("DISCORD_SECRET"), redirectURL, discord.ScopeIdentify))
}

func LoginRouter(ctx context.Context, database mongo.Database) http.Handler {
	// goth.UseProviders(discord.New(os.Getenv("DISCORD_KEY"), os.Getenv("DISCORD_SECRET"), "http://localhost:3000/auth/discord/callback", discord.ScopeIdentify))
	gothSetup()

	routerModel := LoginRouterConfig{
		database: database,
		ctx:      ctx,
	}
	chi := chi.NewRouter()
	chi.Get("/{provider}", routerModel.LoginHandler)
	chi.Get("/{provider}/callback", routerModel.LoginCallback)
	return chi
}

func (configs LoginRouterConfig) LoginHandler(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	// fmt.Printf(req.Context())
	// fmt.Printf(req.Context().Value("provider").(string))
	providerParam := chi.URLParam(req, "provider")
	ctx := context.WithValue(req.Context(), "provider", providerParam)
	// requestHandler(w, r.WithContext(ctx))
	// fmt.Printf("provider %s \n", req.URL.Query().Get("provider"))
	gothUser, err := gothic.CompleteUserAuth(res, req.WithContext(ctx))
	if err != nil {
		log.Printf("LOGIN HANDLER: user not yet logged in: %s \n", gothUser.RawData)
		gothic.BeginAuthHandler(res, req)
	} else {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("logged in"))
		log.Printf("LOGIN HANDLER: User already logged in: %s ", gothUser.RawData)
		return
	}
}

func (configs LoginRouterConfig) LoginCallback(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		fmt.Fprintln(res, err)
		return
	}
	// err = ValidateUserInDatabase(configs.ctx, user.UserID, &configs.database)
	// if err != nil {
	// 	res.WriteHeader(http.StatusInternalServerError)
	// 	res.Write([]byte(err.Error()))
	// 	fmt.Fprintln(res, err)
	// 	return
	// }

	resBody := []byte(user.UserID)

	res.Write(resBody)
	// t, _ := template.New("foo").Parse(userTemplate)
	// t.Execute(res, user)
	// if r.FormValue("state") != state {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("State does not match."))
	// 	return
	// }
	// token, err := configs.oauth2.Exchange(context.Background(), r.FormValue("code"))

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// res, err := configs.oauth2.Client(context.Background(), token).Get("https://discord.com/api/users/@me")
	// if err != nil || res.StatusCode != 200 {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	if err != nil {
	// 		w.Write([]byte(err.Error()))
	// 	} else {
	// 		w.Write([]byte(res.Status))
	// 	}
	// 	return
	// }
	// defer res.Body.Close()

}

// 	http.Redirect(w, r, url, http.StatusFound)
// }

// func loginCallbackHandler(w http.ResponseWriter, r *http.Request, oauthCFG oauth2.Config) {
// 	fmt.Printf("In Login Callback Handler, URL QUERY: \n %s \n", r.URL.Query())
// 	code := r.URL.Query().Get("code") // TODO:  we should analyze this later

// 	if code == "" {
// 		http.Error(w, "Missing authorization code", http.StatusBadRequest)
// 		return
// 	}

// 	token, err := oauthCFG.Exchange(context.Background(), code)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	http.SetCookie(w, &http.Cookie{
// 		Name:  "access_token",
// 		Value: token.AccessToken})
// 	http.Redirect(w, r, "/", http.StatusFound)
// }

/*

Process Flow for Creating a New User
Receive Callback: When the OAuth provider sends back the authorization code in the callback function, you’ll use this code to request an access token.

Exchange Authorization Code for Access Token: After receiving the access_token and optionally a refresh_token, you can use the access_token to fetch user information.

Request User Information: You will typically call the OAuth provider’s user info endpoint (e.g., Google’s https://www.googleapis.com/oauth2/v3/userinfo) to get the user’s profile data.

Create or Update User Record:

New User: If the user doesn’t already exist in your system (based on email or OAuth provider ID), you’ll create a new user record in your database using the retrieved info (name, email, etc.).
Existing User: If the user exists (e.g., based on email or their OAuth provider ID), update any necessary info or just authenticate them.
Token Management: If you need to keep the user logged in to your own app, you’ll typically generate and store your own session tokens (e.g., JWTs) rather than the OAuth provider’s access token.

*/
