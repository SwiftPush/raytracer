package main

import (
	"log"
	"os"
	"raytracer/internal"
)

func main() {
	err := internal.Render()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}
