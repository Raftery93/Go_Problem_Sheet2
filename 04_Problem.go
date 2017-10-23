// Problem 1
// Conor Raftery 12/10/17

package main

import (
	"html/template"//add html/template package 
	"net/http"
	"math/rand"//imports math random package
	"time"
	"strconv"
)

type myMsg struct {
    Message string
}

//msg1 := message{msg: "Guess a number between 1 and 20"}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	//Browser renders html tags
	w.Header().Set("Content-Type","text/html")

	//Output to browser

	http.ServeFile(w, r, "04_index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request){

	//http.ServeFile(w, r, "04_guess.html")

		//create and initialise string
		message :="Guess a number between 1 and 20"
		
			//set up a seed for random number generator
			//https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly
			rand.Seed(time.Now().UTC().UnixNano())
		
			target:=0//added to delete undefined issue
			var cookie, err = r.Cookie("target")//gets cookie called count
		
			if err == nil{
				//if we could read it ,try to convert its value to an int
				target, _ = strconv.Atoi(cookie.Value)
				if target ==0{
					target = rand.Intn(20-1)
				}
			}
		
			//set cookie details at pointer address to &http.cookie
			cookie = &http.Cookie{
				Name: "target",
				Value: strconv.Itoa(target),
				Expires: time.Now().Add(72 * time.Hour),
			}
			
			//set the cookie
			http.SetCookie(w,cookie)
			
			//read the contents of guess.html and return a template
			t, _ := template.ParseFiles("04_guess.tmpl")
			
			//execute template and pass pointer to myMsg struct
			t.Execute(w, &myMsg{Message:message})
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}