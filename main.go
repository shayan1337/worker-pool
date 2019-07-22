package main

import (
	"log"
	. "worker-pool/url-fetcher"
)

func main() {
	response, err := Fetch("http://facebook.com")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
}
