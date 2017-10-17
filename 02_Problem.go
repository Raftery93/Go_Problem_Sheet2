// Problem 1
// Conor Raftery 12/10/17

package main

import (
	"fmt"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Browser renders html tags
	w.Header().Set("Content-Type","text/html")

	//Output to browser
	fmt.Fprintln(w, "<h1>Guessing Game</h1>")

}

func main() {
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}