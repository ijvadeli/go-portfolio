package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	// Define useable files
	files := []string{
		"./static/html/base.tmpl",
		"./static/html/pages/home.tmpl",
		"./static/html/components/nav.tmpl",
		"./static/html/components/head.tmpl",
	}

	// Error handling for when file is not found
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Error handling for when specific template (in this case base) is not found
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	// Set up mux
	mux := http.NewServeMux()

	// Fileserver for all static files in assets
	fs := http.FileServer(http.Dir("./static/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets", neuter(fs)))

	// Page paths
	mux.HandleFunc("GET /{$}", home) // {$} Means path must be exactly with "/"

	// Setting up host
	port := ":8080"
	log.Print("Starting on port :8080")
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}

// Custom middleware to intercept requests before they reach FileServer
func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}
