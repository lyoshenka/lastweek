package main

import (
	"fmt"
	"github.com/zenazn/goji/web"
	"net/http"
	"net/url"
	"time"

	gocache "github.com/pmylund/go-cache"
)

type RouteMap map[string]bool

//
// AUTH MIDDLEWARE
//

func SessionMiddleware(e *Env, unauthRoutes RouteMap) func(*web.C, http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			e.Session = getSession(w, r)
			if e.Session.GithubToken == "" && !unauthRoutes[r.URL.Path] {
				state := randString(16)
				e.Session.GithubAuthState = state
				queryParams := url.Values{}
				queryParams.Add("client_id", e.GithubClientId)
				queryParams.Add("scope", "repo")
				queryParams.Add("state", state)
				http.Redirect(w, r, fmt.Sprintf("https://github.com/login/oauth/authorize?%s", queryParams.Encode()), http.StatusFound)
				return
			}
			h.ServeHTTP(w, r)
		})
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
