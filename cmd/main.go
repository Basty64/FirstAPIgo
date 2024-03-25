package main

import (
	"MyFirstAPIgo/internal/entities"

	"context"

	"log"
)

func main() {

	ctx := context.Background()
	server, err := entities.NewServer(ctx, "localhost:8080", "")
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("error: %s", err)
	}

	err := server.Shutdown(ctx)
	if err != nil {
		return
	}
}

