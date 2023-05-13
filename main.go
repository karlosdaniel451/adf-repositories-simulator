package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	port int = 8000
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/api/adfweb-version/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(GetAvailableADFWebVersions())
	})

	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("Starting server on:%d\n", port)
	log.Fatal(server.ListenAndServe())
}

func GetAvailableADFWebVersions() []string {
	availableVersions := []string{}

	// TODO: Implement the logic to check the available versions localised in `static/adfweb/`
	availableVersions = append(availableVersions, "0.0.1")
	availableVersions = append(availableVersions, "0.0.2")
	availableVersions = append(availableVersions, "0.0.3")

	return availableVersions
}
