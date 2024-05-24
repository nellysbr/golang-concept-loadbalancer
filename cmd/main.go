package main

import (
	"fmt"
	"time"
)

func worker (id int, jobs chan int) {
	for j := range jobs {
		fmt.Printf("Worker %d received %d\n", id, j)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int, 5)

	// Start 5 workers
	for w := 1; w <= 5; w++ {
		go worker(w, jobs)
	}
	for j := 1; j <= 100; j++ {
		jobs <- j
	}
	close(jobs)
	
}