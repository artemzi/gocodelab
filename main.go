package main

import (
	"flag"
	"log"

	"github.com/artemzi/gocodelab/api"
)

func main() {
	bindAddr := flag.String("bind_addr", ":8080", "Set bind address")
	flag.Parse()
	a := api.New(*bindAddr)
	log.Fatal(a.Start())
}
