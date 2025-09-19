package main

import (
	"log"

	"example.com/gin-api/internal/album"
	"example.com/gin-api/internal/router"
)

func main() {
	r := router.New()
	album.NewHandler().Register(r)

	addr := "localhost:8080"
	log.Printf("listening on http://%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}
