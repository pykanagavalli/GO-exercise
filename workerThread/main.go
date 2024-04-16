package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("Worker", id, "Started", j)
		time.Sleep(time.Second)
		fmt.Println("Worker", id, "Finished", j)
		results <- j * 2
	}
}

func main() {
	const numjobs = 5
	jobs := make(chan int, numjobs)
	results := make(chan int, numjobs)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= numjobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numjobs; a++ {
		<-results
	}
}
