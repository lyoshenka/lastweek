package main

import (
	"os"

	"github.com/elazarl/goproxy"

	"github.com/zenazn/goji"
)

func main() {

	_ = goproxy.CA_CERT // so goproxy can be included and go-gettable. its not if you don't include it explicitly

	env := &Env{
		GithubClientId:     os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	}

	goji.Use(SessionMiddleware(env))

	goji.Get("/git_auth_hook", Handler{env, gitAuthRoute})
	goji.Get("/commits", Handler{env, commitsRoute})

	goji.Use(StaticMiddleware(env, "static"))

	goji.Serve()
}

//
// UTIL
//

func check(e error) {
	if e != nil {
		panic(e)
	}
}
