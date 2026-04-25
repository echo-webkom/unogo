package unogo

import (
	"errors"
	"net/http"
)

const AuthCallbackEndpoint = "/api/auth/callback"
const AuthErrorEndpoint = "/api/logg-inn"

type LoginCallback func(w http.ResponseWriter, r *http.Request, sessionToken string)
type ErrorCallback func(w http.ResponseWriter, r *http.Request, err error, attemptId string)

type GenericMux interface {
	Handle(path string, handler http.Handler)
}

// MountLogin will mount both the LoginHandler and LoginErrorHandler to the correct endpoints and
// pass the given login and error callbacks. This function can be used if your mux/router implements
// the GenericMux interface (for example http.Mux and chi.Mux/Router).
func MountLogin(mux GenericMux, lc LoginCallback, ec ErrorCallback) {
	mux.Handle(AuthCallbackEndpoint, LoginHandler(lc))
	mux.Handle(AuthErrorEndpoint, LoginErrorHandler(ec))
}

// LoginHandler is a http.Handler you can mount to AuthCallbackEndpoint to handle login requests.
// Note that this handler must be mounted at this exact endpoint as it is what Uno will redirect to
// after a successful login attempt.
//
// Unpon a successful login, [callback] will be called with the users session token.
func LoginHandler(callback LoginCallback) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		callback(w, r, token)
	})
}

// LoginErrorHandler is a http.Handler you can mount to AuthErrorEndpoint to handle failed login
// requests. Note that this handler must be mounted at this exact endpoint as it is what Uno will
// redirect to after an unsuccessful login attempt.
//
// Unpon a failed login, [callback] will be called with the error message and login attempt id.
func LoginErrorHandler(callback ErrorCallback) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.URL.Query().Get("token")
		attemptId := r.URL.Query().Get("token")
		callback(w, r, errors.New(err), attemptId)
	})
}
