package main

import (
	"YoutubeBingo/internal/handlers"
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", ":8080", "Port to listen on")
	flag.Parse()

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/bingo", handlers.GetRandomBingoItem)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Printf("Server running on http://localhost%s\n", *port)
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
