package main

import (
	"fmt"
	"log"
	"net/http"

	backend "will-steinleitner.de"
	"will-steinleitner.de/internal/server"
)

func main() {
	fmt.Println("Starting server at port 8090...")
	app := backend.NewApplication()

	mux := server.RegisterRoutes(app.GetHome())
	if err := http.ListenAndServe(":8090", mux); err != nil {
		log.Fatal(err)
	}
}
