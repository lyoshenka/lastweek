package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/parnurzeal/gorequest"
	"github.com/zenazn/goji/web"
)

func homeRoute(e *Env, c web.C, w http.ResponseWriter, r *http.Request) error {
	if e.Session.GithubToken == "" {
		state := randString(16)
		e.Session.GithubAuthState = state
		queryParams := url.Values{}
		queryParams.Add("client_id", e.GithubClientId)
		queryParams.Add("scope", "repo")
		queryParams.Add("state", state)
		http.Redirect(w, r, fmt.Sprintf("https://github.com/login/oauth/authorize?%s", queryParams.Encode()), http.StatusFound)
		return nil
	}

	templateArgs := map[string]interface{}{
		"token": e.Session.GithubToken,
	}
	fmt.Fprintln(w, getTemplate("home", templateArgs))
	return nil
}

func robotsRoute(e *Env, c web.C, w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "User-agent: *\nDisallow: /")
	return nil
}

func commitsRoute(e *Env, c web.C, w http.ResponseWriter, r *http.Request) error {
	if e.Session.GithubToken == "" {
		return StatusError{401, fmt.Errorf("Please authenticate with Github")}
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	_, body, _ := gorequest.New().Get("https://api.github.com/repos/topscore/topscore/commits").
		Param("access_token", e.Session.GithubToken).
		Param("sha", "master").
		Param("page", page).
		Param("per_page", "100").
		End()

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(body))
	return nil
}

func gitAuthRoute(e *Env, c web.C, w http.ResponseWriter, r *http.Request) error {
	code := r.URL.Query().Get("code")
	if code == "" {
		return StatusError{400, fmt.Errorf("No 'code' query param")}
	}
	state := r.URL.Query().Get("state")
	if state == "" {
		return StatusError{400, fmt.Errorf("No 'state' query param")}
	}
	if state != e.Session.GithubAuthState {
		return StatusError{400, fmt.Errorf("Query state doesnt match saved state")}
	}

	data := url.Values{}
	data.Add("client_id", e.GithubClientId)
	data.Add("client_secret", e.GithubClientSecret)
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

	e.Session.GithubToken = responseData.Get("access_token")
	e.Session.GithubAuthState = ""
	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}
