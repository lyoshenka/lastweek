package main

import (
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/zenazn/goji/web"
)

func SessionMiddleware(e *Env) func(*web.C, http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			e.Session = getSession(w, r)
			h.ServeHTTP(w, r)
		})
	}
}

func StaticMiddleware(e *Env, directory string) func(*web.C, http.Handler) http.Handler {
	dir := http.Dir(directory)
	indexFile := "index.html"
	logStaticFilesServed := false
	return func(c *web.C, h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "HEAD" && r.Method != "GET" {
				h.ServeHTTP(w, r)
				return
			}

			if r.URL.Path == "/" && requireAuth(e, w, r) {
				return
			}

			// Get the file name from the path
			file := r.URL.Path

			// Open the file and get the stats
			f, err := dir.Open(file)
			if err != nil {
				h.ServeHTTP(w, r)
				return
			}
			defer f.Close()

			fs, err := f.Stat()
			if err != nil {
				h.ServeHTTP(w, r)
				return
			}

			// if the requested resource is a directory, try to serve the index file
			if fs.IsDir() {
				// redirect if trailling "/"" is missing
				if !strings.HasSuffix(r.URL.Path, "/") {
					http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
					return
				}

				file = path.Join(file, indexFile)
				f, err = dir.Open(file)
				if err != nil {
					h.ServeHTTP(w, r)
					return
				}
				defer f.Close()
				fs, err = f.Stat()
				if err != nil || fs.IsDir() {
					h.ServeHTTP(w, r)
					return
				}
			}

			if logStaticFilesServed {
				log.Println("[Static] Serving " + file)
			}

			// Add an Expires header to the static content
			// if opt.Expires != nil {
			// 	w.Header().Set("Expires", opt.Expires())
			// }

			http.ServeContent(w, r, file, fs.ModTime(), f)

			h.ServeHTTP(w, r)
		})
	}
}
