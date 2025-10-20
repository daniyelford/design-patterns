package main

import "fmt"

func iterator(slice []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, v := range slice {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func main() {
	for v := range iterator([]int{1, 2, 3}) {
		fmt.Println(v)
	}
}
