package main

import (
	"fmt"
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

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./public/about.html")
	if err != nil {
		fmt.Print(err)
	}
	w.Write([]byte(b))
}

func linuxHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./public/linux.html")
	if err != nil {
		fmt.Print(err)
	}
	w.Write([]byte(b))
}

func appleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./public/apple.html")
	if err != nil {
		fmt.Print(err)
	}
	w.Write([]byte(b))
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("./public/projects.html")
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

	about := http.HandlerFunc(aboutHandler)
	http.HandleFunc("/about", about)

	linux := http.HandlerFunc(linuxHandler)
	http.HandleFunc("/linux", linux)

	apple := http.HandlerFunc(appleHandler)
	http.HandleFunc("/apple", apple)

	projects := http.HandlerFunc(projectsHandler)
	http.HandleFunc("/projects", projects)

	fmt.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
