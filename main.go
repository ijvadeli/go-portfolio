package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./static"))
	fs.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", home)

	port := ":8080"
	fmt.Println("Server running on " + port)

	log.Fatal(http.ListenAndServe(port, nil))
}
