package main

import "fmt"

func worker(id int, jobs <-chan int, out chan<- int) {
	for j := range jobs {
		out <- j * id
	}
}

func main() {
	jobs := make(chan int, 3)
	results := make(chan int, 9)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 3; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 0; i < 9; i++ {
		fmt.Println(<-results)
	}
}
