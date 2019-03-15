package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to Site</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:dfox@pmh-mb.ca\">support</a>.")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>This is not the page you are looking for :(</h1><p>If you see this lots, TOO BAD!!!</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
