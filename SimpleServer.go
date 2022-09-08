 HTTP server in Go

// package contains all utilities needed to accept requests and handle them dynamically.
package main

import (
    "fmt"
    "net/http" HandleFunc registers the handler function for the given pattern in the DefaultServeMux
)

func main() {
// package contains all utilities needed to accept requests and handle them dynamically.
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
// ile server to work properly it needs to know, where to serve files from
        fmt.Fprintf(w, "Welcome, We are Up!") // Prints out messgae
    })
	// handles storage and management of data files and accessibility
    fs := http.FileServer(http.Dir("static/"))
// In order to serve files correctly, we need to strip away a part of the url path.
    http.Handle("/static/", http.StripPrefix("/static/", fs))
 //the function is use to initialize our server
    http.ListenAndServe(":8080", nil)
}
