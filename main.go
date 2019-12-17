package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("PDTOKEN")
	schedules := os.Getenv("PDSHEDULES")
	pd := NewPdApi(token, schedules)
	fs := http.FileServer(http.Dir("./public/static"))
	handlers := NewHandlers(pd)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", handlers)
	if os.Getenv("PDVIEWERLISTEN") == "" {
		log.Println("PDVIEWERLISTEN not set. Using  0.0.0.0:9090")
		http.ListenAndServe(":9090", nil)
	} else {
		log.Println("Listening on "+os.Getenv("PDVIEWERLISTEN"))
		http.ListenAndServe(os.Getenv("PDVIEWERLISTEN"), nil)
	}
}
