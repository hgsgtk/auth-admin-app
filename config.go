package main

import (
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	OAuthConfig *oauth2.Config

	SessionStore sessions.Store
)

func init() {
	// auth
	OAuthConfig = configureOAuthClient("client-id", "client-key")

	// sessions
	// cookie store
	cookieStore := sessions.NewCookieStore([]byte("secret-key"))
	cookieStore.Options = &sessions.Options{
		HttpOnly: true,
	}
	SessionStore = cookieStore

	// FIXME: Redis
	// store, err := redistore.NewRediStore(10, "tcp", ":6379", "", []byte("Ad8vC0Kl5O81SoM1EjnZ88EAlNBd9XxTpsg43DEk"))
	// if err != nil {
	// 	panic(err)
	// }
	// defer store.Close()
	// SessionStore = store
}

func configureOAuthClient(clientID, clientSecret string) *oauth2.Config {
	redirectURL := os.Getenv("OAUTH2_CALLBACK")
	if redirectURL == "" {
		redirectURL = "http://localhost:8080/oauth2callback"
	}
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
}
