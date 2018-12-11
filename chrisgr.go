package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func handleError(info string, err error) {
	if err != nil {
		fmt.Println("info: ", info)
		panic(err)
	}
}

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

func templatePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/template.html"))

	type TmplPageData struct {
		Links template.HTML
		Body  template.HTML
	}

	log.Println("static page template ", r.URL.Path)
	err := r.ParseForm()
	handleError("read template file", err)

	file := r.URL.Path

	if file == "/" {
		file = "/index"
	}

	if _, err := os.Stat("public" + file + ".html"); !os.IsNotExist(err) {

		bodyString, rferr := ioutil.ReadFile("public" + file + ".html")
		handleError("read page file", rferr)

		headerLinksString, rferr2 := ioutil.ReadFile("public/header_links.html")
		handleError("read header file", rferr2)

		headerLinks := template.HTML(string(headerLinksString))
		body := template.HTML(string(bodyString))

		data := TmplPageData{
			Links: headerLinks,
			Body:  body,
		}

		tmpl_err := tmpl.Execute(w, data)
		handleError("template execute error:", tmpl_err)
	} else {
		fmt.Fprintf(w, "404")
	}
}

func contactFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("public/send_contact.html"))

	headerLinksString, rferr2 := ioutil.ReadFile("public/header_links.html")
	handleError("read header file", rferr2)
	headerLinks := template.HTML(string(headerLinksString))

	type TmplPageData struct {
		Links   template.HTML
		Name    string
		Email   string
		Subject string
		Message string
	}

	log.Println("Contact form handler ", r.URL.Path)
	err := r.ParseForm()
	handleError("parse form error:", err)

	name := r.Form.Get("name")
	email := r.Form.Get("email")
	subject := r.Form.Get("subject")
	message := r.Form.Get("message")

	data := TmplPageData{
		Links:   headerLinks,
		Name:    name,
		Email:   email,
		Subject: subject,
		Message: message,
	}

	tmpl_err := tmpl.Execute(w, data)
	handleError("templ execute error:", tmpl_err)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("static handler: ", r.URL.Path)
	b, err := ioutil.ReadFile("./public" + r.URL.Path + ".html")
	if err != nil {
		fmt.Print(err)
	}
	w.Write([]byte(b))
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	health := http.HandlerFunc(healthHandler)
	http.HandleFunc("/health", health)

	th := http.HandlerFunc(timeHandler)
	http.HandleFunc("/time", th)

	root := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/", root)

	index := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/index", index)

	about := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/about", about)

	linux := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/linux", linux)

	apple := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/apple", apple)

	projects := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/projects", projects)

	experiments := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/experiments", experiments)

	page := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/template", page)

	contact := http.HandlerFunc(templatePageHandler)
	http.HandleFunc("/contact", contact)

	send_contact := http.HandlerFunc(contactFormHandler)
	http.HandleFunc("/send_contact", send_contact)

	fmt.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
