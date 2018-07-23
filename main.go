package main

import (
	"log"

	"github.com/artemzi/gocodelab/api"
)

func main() {
	a := api.New(":9111")
	log.Fatal(a.Start())
}
