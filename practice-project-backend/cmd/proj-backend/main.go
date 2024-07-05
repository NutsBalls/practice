package main

import (
	"log"

	"github.com/NutsBalls/practice-project-backend/internal/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run(8080)
	if err != nil {
		log.Fatal(err)
	}
}
