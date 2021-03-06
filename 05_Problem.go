// Problem 1
// Conor Raftery 12/10/17

package main

import (
	"html/template"
	"net/http"
	"math/rand"
	"time"
	"strconv"
)

type myMsg struct {
    Message string
}

//msg1 := message{msg: "Guess a number between 1 and 20"}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","text/html")

	http.ServeFile(w, r, "05_Problem.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request){

	//http.ServeFile(w, r, "04_guess.html")

		message :="Guess a number between 1 and 20"
		
			rand.Seed(time.Now().UTC().UnixNano())
		
			target:= rand.Intn(20-1)
			
			var cookie, err = r.Cookie("target")
		
			if err == nil{
				
				target, _ = strconv.Atoi(cookie.Value)

			}
		
			cookie = &http.Cookie{
				Name: "target",
				Value: strconv.Itoa(target),
				Expires: time.Now().Add(72 * time.Hour),
			}
			
			http.SetCookie(w,cookie)
			
			t, _ := template.ParseFiles("05_Problem.tmpl")

			t.Execute(w, &myMsg{Message:message})
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}