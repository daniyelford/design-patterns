package main

import "fmt"

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Println("Produced", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println("Consumed", v)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
