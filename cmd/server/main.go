package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/sinmetal/ff14cf"
	metadatabox "github.com/sinmetalcraft/gcpbox/metadata"
)

func main() {
	ctx := context.Background()

	log.Print("starting server...")

	pID, err := metadatabox.ProjectID()
	if err != nil {
		log.Fatal(err)
	}
	fs, err := firestore.NewClient(ctx, pID)
	if err != nil {
		log.Fatal(err)
	}

	stores := &Stores{}

	userStore, err := ff14cf.NewUserStore(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	stores.UserStore = userStore

	contentStore, err := ff14cf.NewContentStore(ctx, fs)
	if err != nil {
		log.Fatal(err)
	}
	stores.ContentStore = contentStore

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

	registerHandler, err := ff14cf.NewRegisterHandler(ctx, userStore)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/api/register", registerHandler.Handle)

	contentHandler, err := ff14cf.NewContentsHandler(ctx, stores.ContentStore)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/content/", contentHandler.Handle)

	contentsQuizHandler, err := ff14cf.NewContentsQuizHandler(ctx, stores.ContentStore)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/api/content/quiz", contentsQuizHandler.Handle)

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

type Stores struct {
	UserStore    *ff14cf.UserStore
	ContentStore *ff14cf.ContentStore
}
