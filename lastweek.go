//go:generate go-bindata -pkg $GOPACKAGE -o static.go templates/

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/lyoshenka/go-bindata-html-template"
	baseTemplate "html/template"

	"github.com/elazarl/goproxy"
	"github.com/parnurzeal/gorequest"

	gocache "github.com/pmylund/go-cache"

	"github.com/hypebeast/gojistatic"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

const (
	cookieName         = "lwsession"
	githubClientId     = "61f99518daf157383594"
	githubClientSecret = "b83af5782f534917967911263c2473871b962fb9"
)

func main() {

	_ = goproxy.CA_CERT // so goproxy can be included and go-gettable. its not if you don't include it explicitly

	err := loadTemplates()
	check(err)

	goji.Use(gojistatic.Static("static", gojistatic.StaticOptions{SkipLogging: true}))

	goji.Get("/", homeRoute)
	goji.Get("/git_auth_hook", gitAuthRoute)
	goji.Get("/commits", commitsRoute)
	goji.Get("/robots.txt", robotsRoute)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	check(err)

	goji.ServeListener(listener)
}

func homeRoute(c web.C, w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}

	session := getSession(w, r)

	if session.GithubToken == "" {
		state := randString(16)
		session.GithubAuthState = state
		queryParams := url.Values{}
		queryParams.Add("client_id", githubClientId)
		queryParams.Add("scope", "repo")
		queryParams.Add("state", state)
		queryParams.Add("redirect_url", "http://localhost:8000/git_auth_hook")
		http.Redirect(w, r, fmt.Sprintf("https://github.com/login/oauth/authorize?%s", queryParams.Encode()), http.StatusFound)
		return
	}

	templateArgs := map[string]interface{}{
		"token": session.GithubToken,
	}
	fmt.Fprintln(w, getTemplate("home", templateArgs))
}

func robotsRoute(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "User-agent: *\nDisallow: /")
}

func commitsRoute(c web.C, w http.ResponseWriter, r *http.Request) {
	session := getSession(w, r)

	if session.GithubToken == "" {
		fmt.Fprintln(w, "Not authenticated")
		return
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	_, body, _ := gorequest.New().Get("https://api.github.com/repos/topscore/topscore/commits").
		Param("access_token", session.GithubToken).
		Param("sha", "master").
		Param("page", page).
		Param("per_page", "100").
		End()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
}

func gitAuthRoute(c web.C, w http.ResponseWriter, r *http.Request) {
	session := getSession(w, r)

	code := r.URL.Query().Get("code")
	if code == "" {
		fmt.Fprintln(w, "No 'code' query param")
		return
	}
	state := r.URL.Query().Get("state")
	if state == "" {
		fmt.Fprintln(w, "No 'state' query param")
		return
	}
	if state != session.GithubAuthState {
		fmt.Fprintln(w, "Query state doesnt match saved state")
		return
	}

	data := url.Values{}
	data.Add("client_id", githubClientId)
	data.Add("client_secret", githubClientSecret)
	data.Add("code", code)
	data.Add("state", state)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	responseData, err := url.ParseQuery(string(body))
	check(err)

	session.GithubToken = responseData.Get("access_token")
	session.GithubAuthState = ""
	http.Redirect(w, r, "/", http.StatusFound)
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
// RAND
//

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var randSrc = rand.NewSource(time.Now().UnixNano())

func randString(n int) string {
	b := make([]byte, n)
	// A randSrc.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

//
// UTIL
//

func check(e error) {
	if e != nil {
		panic(e)
	}
}
