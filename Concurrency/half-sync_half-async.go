package main

import (
	"fmt"
	"time"
)

func asyncReceiver(jobs chan int, results chan int) {
	for job := range jobs {
		go func(j int) {
			time.Sleep(500 * time.Millisecond)
			results <- j * 2 // async
		}(job)
	}
}

func syncProcessor(results chan int) {
	for res := range results {
		fmt.Println("Processed:", res) // sync
	}
}

func main() {
	jobs := make(chan int, 3)
	results := make(chan int, 3)

	go asyncReceiver(jobs, results)
	go syncProcessor(results)

	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	time.Sleep(2 * time.Second)
}
