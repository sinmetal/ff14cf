package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sinmetal/ff14cf"
)

func main() {
	ctx := context.Background()

	log.Print("starting server...")

	staticContentsHandler, err := ff14cf.NewStaticContentsHandler(ctx)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/register.html", staticContentsHandler.Handler)

	loadStoneHandler, err := ff14cf.NewLoadStoneHandler(ctx)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/api/lodestone", loadStoneHandler.Handle)

	http.HandleFunc("/", handler)
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
