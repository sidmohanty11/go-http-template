package main

import (
	"flag"
	"fmt"
	"log"
	"mygoapp/api"
	"mygoapp/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "the server address")
	flag.Parse()

	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("Server running at PORT", *listenAddr)
	log.Fatal(server.Start())
}
