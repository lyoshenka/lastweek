//go:generate go-bindata -pkg $GOPACKAGE -o static.go templates/

package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/lyoshenka/go-bindata-html-template"
	baseTemplate "html/template"

	"github.com/elazarl/goproxy"

	"github.com/hypebeast/gojistatic"
	"github.com/zenazn/goji"
)

func main() {

	_ = goproxy.CA_CERT // so goproxy can be included and go-gettable. its not if you don't include it explicitly

	err := loadTemplates()
	check(err)

	env := &Env{
		GithubClientId:     os.Getenv("GITHUB_CLIENT_ID"),
		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	}

	goji.Use(gojistatic.Static("static", gojistatic.StaticOptions{SkipLogging: true}))

	goji.Get("/", Handler{env, homeRoute})
	goji.Get("/git_auth_hook", Handler{env, gitAuthRoute})
	goji.Get("/commits", Handler{env, commitsRoute})
	goji.Get("/robots.txt", Handler{env, robotsRoute})

	goji.Serve()
}

//
// TEMPLATES
//

var templates *template.Template

func loadTemplates() error {
	t, err := template.New("mytmpl", Asset).Funcs(template.FuncMap{
		"safehtml": func(value interface{}) baseTemplate.HTML {
			return baseTemplate.HTML(fmt.Sprint(value))
		},
	}).ParseFiles(AssetNames()...)

	if err == nil {
		templates = t
	}
	return err
}

func getTemplate(name string, data interface{}) string {
	var doc bytes.Buffer
	templates.ExecuteTemplate(&doc, name, data)
	return doc.String()
}

//
// UTIL
//

func check(e error) {
	if e != nil {
		panic(e)
	}
}
