package main

import (
	"log"
	. "worker-pool/job"
	. "worker-pool/url-fetcher"
)

func main() {
	job := NewJob(1, "http://facebook.com")
	result := Fetch(job)
	if result.Err() != nil {
		log.Fatal(result.Err())
	}
	log.Println(result.Response())
}
