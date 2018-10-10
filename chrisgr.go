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

	fmt.Println("Server: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
