package main

import (
	"net/http"
	"time"

	gocache "github.com/pmylund/go-cache"
)

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
