package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	//"html/template"
)

const (
	username = "admin"
	password = "password123"
	authorizationKey = "my-secret-token"
)

func hashCredentials(username, password string) string {
	hash := sha256.New()
	hash.Write([]byte(username + ":" + password))
	return hex.EncodeToString(hash.Sum(nil))
}

func authenticate(token string) bool {
	return token == hashCredentials(username, password)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	if r.Form.Get("username") == username && r.Form.Get("password") == password {
		token := hashCredentials(username, password)
		//fmt.Println("Generated Token:", token) // Print the generated token
		http.Redirect(w, r, "/form.html?token="+token, http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, "Invalid username or password!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//token := r.URL.Query().Get("token")

	//tmpl, err := template.ParseFiles("form.html")

	/*if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}*/

	/*err = tmpl.Execute(w, struct{ Token string }{Token: token})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}*/

	//fmt.Println("Received Token:", token) // Print the received token
	/*if !authenticate(token) {
		fmt.Println("Authentication failed for token:", token)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}*/

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	name := r.Form.Get("name")
	message := r.Form.Get("message")

	fmt.Fprintf(w, "POST request successful\n")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Message = %s\n", message)
}



func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Wah Amjinc Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", HelloHandler)

	fmt.Printf("Starting at port 8080 || ")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}