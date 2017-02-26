package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {

	filteredRequest := map[string]interface{}{
		"method":     r.Method,
		"url":        r.URL.Path,
		"proto":      r.Proto,
		"header":     r.Header,
		"host":       r.Host,
		"remoteAddr": r.RemoteAddr,
		"requestURI": r.RequestURI,
	}

	b, err := json.MarshalIndent(filteredRequest, "", "  ")
	if err != nil {
		log.Print("Error:", err, "occured")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Sorry a problem occured")

	} else {
		raw, _ := json.Marshal(filteredRequest)
		log.Print(string(raw))

		formatted := string(b)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, formatted)
	}
}

func main() {

	var port string

	if configuredPort := os.Getenv("PORT"); configuredPort != "" {
		port = configuredPort
	} else {
		port = "8080"
	}

	log.Print("Binding service to: ", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
