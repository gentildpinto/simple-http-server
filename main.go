package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/form", formHandler)

	fmt.Printf("Starting the server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s [%s] %s %s\n", r.Host, r.Method, r.Proto, r.URL.Path)
	if r.URL.Path != "/" {
		log.Printf("status=%d message=%s\n", http.StatusNotFound, http.StatusText(http.StatusNotFound))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	files := []string{
		"./templates/base.tmpl.html",
		"./templates/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("status=%d message=%s\n", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("status=%d message=%s\n", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Printf("status=%d message=%s\n", http.StatusOK, http.StatusText(http.StatusOK))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s [%s] %s %s\n", r.Host, r.Method, r.Proto, r.URL.Path)

	files := []string{
		"./templates/base.tmpl.html",
		"./templates/pages/hello.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("status=%d message=%s\n", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("status=%d message=%s\n", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Printf("status=%d message=%s\n", http.StatusOK, http.StatusText(http.StatusOK))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s [%s] %s %s\n", r.Host, r.Method, r.Proto, r.URL.Path)

	files := []string{
		"./templates/base.tmpl.html",
		"./templates/pages/form.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("status=%d message=%s\n", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Printf("status=%d message=%s\n", http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Printf("status=%d message=%s\n", http.StatusOK, http.StatusText(http.StatusOK))
}
