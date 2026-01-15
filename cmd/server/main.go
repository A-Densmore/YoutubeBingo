package main

import (
	"YoutubeBingo/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/bingo", handlers.GetRandomBingoItem)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../web/static"))))

	http.ListenAndServe(":8080", nil)
}
