package main

import (
	"MyFirstAPIgo/internal/entities"
	"log"
)

func main() {

	server := new(entities.Server)
	if err := server.Run("localhost:8080"); err != nil {
		log.Fatalf("error: %s", err)
	}

}
