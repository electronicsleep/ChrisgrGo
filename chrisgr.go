package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	healthStatusResponse := "ok"
	log.Println("Health endpoint")
	fmt.Fprintf(w, "chrisgr: health %s", healthStatusResponse)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Time endpoint")
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("time: " + tm))
}

func contactformHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/template.html"))

	type TmplPageData struct {
		Name    string
		Email   string
		Subject string
		Message string
	}

	log.Println("Contact form endpoint")
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	subject := r.Form.Get("subject")
	message := r.Form.Get("message")

	data := TmplPageData{
		Name:    name,
		Email:   email,
		Subject: subject,
		Message: message,
	}

	tmpl_err := tmpl.Execute(w, data)
	if tmpl_err != nil {
		panic(tmpl_err)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	b, err := ioutil.ReadFile("./public" + r.URL.Path + ".html")
	if err != nil {
		fmt.Print(err)
	}
	w.Write([]byte(b))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	root := http.HandlerFunc(healthHandler)
	http.HandleFunc("/health", root)

	th := http.HandlerFunc(timeHandler)
	http.HandleFunc("/time", th)

	http.Handle("/", http.FileServer(http.Dir("./public/")))

	about := http.HandlerFunc(staticHandler)
	http.HandleFunc("/about", about)

	linux := http.HandlerFunc(staticHandler)
	http.HandleFunc("/linux", linux)

	apple := http.HandlerFunc(staticHandler)
	http.HandleFunc("/apple", apple)

	projects := http.HandlerFunc(staticHandler)
	http.HandleFunc("/projects", projects)

	contact := http.HandlerFunc(staticHandler)
	http.HandleFunc("/contact", contact)

	send_contact := http.HandlerFunc(contactformHandler)
	http.HandleFunc("/send_contact", send_contact)

	fmt.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
