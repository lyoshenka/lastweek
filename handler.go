package main

import (
	"github.com/zenazn/goji/web"
	"log"
	"net/http"
	"time"

	gocache "github.com/pmylund/go-cache"
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	Status() int
	error
}

// StatusError represents an error with an associated HTTP status code.
type StatusError struct {
	Code int
	Err  error
}

// Allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// A (simple) example of our application-wide configuration.
type Env struct {
	Session            *sessionType
	GithubClientId     string
	GithubClientSecret string
}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Env
	H func(e *Env, c web.C, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTPC(c web.C, w http.ResponseWriter, r *http.Request) {
	h.Env.Session = getSession(w, r)
	err := h.H(h.Env, c, w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

//
// SESSION + CACHE
//

var cache = gocache.New(1*time.Hour, 5*time.Minute)

type sessionType struct {
	GithubToken     string
	GithubAuthState string
}

func getSession(w http.ResponseWriter, r *http.Request) *sessionType {
	cookieName := "lwsession"
	cookie, err := r.Cookie(cookieName)
	if err != nil && err != http.ErrNoCookie {
		panic(err)
	}

	if cookie == nil || cookie.Value == "" {
		cookie = &http.Cookie{Name: cookieName, Value: randString(32), Expires: time.Now().Add(time.Hour), HttpOnly: true}
		http.SetCookie(w, cookie)
	}

	session := &sessionType{}
	if cachedVal, found := cache.Get(cookie.Value); found {
		session = cachedVal.(*sessionType)
	} else {
		cache.Set(cookie.Value, session, 0)
	}

	return session
}
