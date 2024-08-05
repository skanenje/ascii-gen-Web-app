package main

import (
	"fmt"
	"net/http"

	"flags/handlers"
)

func main() {
	// Serve static files from the "static" directory
	staticFs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFs))

	imagesFs := http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", imagesFs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)
	http.Handle("/about.html", http.RedirectHandler("/about", http.StatusMovedPermanently))
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/400", handlers.Error400Handler)
	http.HandleFunc("/404", handlers.Error404Handler)
	http.HandleFunc("/405", handlers.Error405Handler)
	http.HandleFunc("/500", handlers.Error500Handler)

	fmt.Println("Starting server at :8080")
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
