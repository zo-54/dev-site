package main

import (
	"log"
	"os"

	"github.com/zo-54/dev-site/backend/server"
)

func main() {
	os.Exit(run())
}

func run() int {
	srv := server.NewServer()

	err := srv.Run()

	if err != nil {
		log.Printf("Failed to start server, err: %v", err)
		return 1
	}

	return 0
}
