package main

import (
	"log"

	"github.com/kacperkg/vellum/internal/server"
)

func main() {
    app, err := server.New()
    if err != nil {
        log.Fatal(err)
    }

    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}