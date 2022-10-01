package main

import (
	"fmt"
	"github.com/jdhenke/gcloud"
	"net/http"
	"os"
)

func main() {
	externalURI := os.Getenv("EXTERNAL_URI") // e.g. http://localhost:8000
	const redirectPath = "/auth/redirect"
	redirectURI := fmt.Sprintf("%s%s", externalURI, redirectPath)
	oauthClient := os.Getenv("OAUTH_CLIENT_ID")
	oauthSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	auth := gcloud.NewAuthorizer(oauthClient, oauthSecret, redirectURI, [][]byte{[]byte("insecure")})
	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.HandleLogin)
	mux.HandleFunc("/logout", auth.HandleLogout)
	mux.HandleFunc(redirectPath, auth.HandleRedirect)
	mux.Handle("/", auth.RequireAuth(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(rw, "Hello", auth.GetSession(r.Context()).Email)
	})))
	panic(http.ListenAndServe(":8000", mux))
}
