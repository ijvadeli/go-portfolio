package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./static/html/base.tmpl",
		"./static/html/pages/home.tmpl",
		"./static/html/components/nav.tmpl",
		"./static/html/components/head.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	// Fileserver for all files in assets to make them reachable
	fs := http.FileServer(http.Dir("./static/assets/"))
	http.Handle("GET /assets/", http.StripPrefix("/assets", fs))

	// Page paths
	http.HandleFunc("/", home)

	// Setting up host
	port := ":8080"
	log.Print("Starting on port :8080")
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}
