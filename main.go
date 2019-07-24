package main

import (
	"fmt"
	"log"
	"sync"
	. "worker-pool/job"
	. "worker-pool/result"
	. "worker-pool/url-fetcher"
)

var jobs = make(chan Job, 5)
var results = make(chan Result, 5)

var urls = []string{"http://facebook.com", "http://instagram.com", "http://linkedin.com", "http://stackoverflow.com", "http://golang.org", "http://helloworld.com", "http://somewebsite.com", "http://ihaverunoutofwebsites.com"}

func main() {
	var wg sync.WaitGroup
	go createWorkers(3)
	wg.Add(1)
	go display(&wg)
	requestId := 1
	for _, url := range urls {
		jobs <- NewJob(requestId, url)
		requestId = requestId + 1
	}
	close(jobs)
	wg.Wait()
}

func createWorkers(numberOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i<numberOfWorkers; i++ {
		wg.Add(1)
		go work(&wg)
	}
	wg.Wait()
	close(results)
}

func work(wg *sync.WaitGroup) {
	for job := range jobs {
		result := Fetch(job)
		results <- NewResult(job, result.Response(), result.Err())
	}
	wg.Done()
}

func display(wg *sync.WaitGroup) {
	for result := range results {fmt.Println("----------------------------")
		fmt.Println(fmt.Sprintf("Response from %s is :::::", result.Job().Url()))
		if result.Err() != nil {
			log.Println(result.Err())
			continue
		}
		log.Println(result.Response())
	}
	wg.Done()
}
