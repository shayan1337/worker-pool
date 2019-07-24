**WIP**


This project involves an implementation of Worker Pools in Golang. Buffered channels are used to implement the input job queue and result queue. Goroutines are used to create a worker pool.

**HOW IT WORKS**


The worker pool is a bunch of goroutines that listen to a buffered channel which contains jobs for the workers. 


A free goroutine picks up a job and executes it and places the result in an output queue which is also a buffered channel. 

A goroutine reads the result from the output queue and prints it to the standard output.

The job consists of a URL that is used by the routine to retrieve response from the URL. The response from the URL is placed on the output queue 


Usage : 
~~~~
go run main.go
  