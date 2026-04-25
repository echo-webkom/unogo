# unogo

Dette er en Go client for Uno som lar deg sette opp nye apper med Uno som backend tjeneste. Uno må være konfigurert til å whiteliste appen du lager for å få tilgang til APIer med autentifisering.

## Bruk

```sh
go get github.com/echo-webkom/unogo
```

## Eksempel

Eksempel på innlogging med Feide og henting av bruker. Etter man besøker `uno.echo-webkom.com/auth/feide?site=...` og fullfører innlogging vil bruker-IDen bli printet ut:

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/echo-webkom/unogo"
)

func main() {
	client := unogo.NewClient(nil)
	mux := http.NewServeMux()

	unogo.MountLogin(mux, loginHandler(client), errorHandler)
	http.ListenAndServe(":8080", mux)
}

func loginHandler(client *unogo.Client) unogo.LoginCallback {
	return func(w http.ResponseWriter, r *http.Request, sessionToken string) {
		user, err := client.GetUser(sessionToken)
		if err != nil {
			log.Println("Failed to get user: " + err.Error())
			return
		}

		log.Println("Login success! UserId: " + user.ID)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, err error, _ string) {
	log.Println("Failed to log in: " + err.Error())
}
```
