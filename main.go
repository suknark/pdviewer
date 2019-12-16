package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("PDTOKEN")
	schedules := os.Getenv("SHEDULES")
	pd := NewPdApi(token, schedules)
	fs := http.FileServer(http.Dir("./public/static"))
	handlers := NewHandlers(pd)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", handlers)
	log.Println("Listening...")
	http.ListenAndServe(os.Getenv("HOST"), nil)
}
