package main

// import the packages 
// fmt: for print the text
// net/http: to handle a request and response and start server.
import (
	"fmt"
	"net/http"
)

// main(): main function is the entry point in golang program 
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // default route handling in golang 
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path) // print the text
	})
	http.ListenAndServe(":80", nil) // start server
}