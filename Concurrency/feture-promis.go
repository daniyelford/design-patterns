package main

import "fmt"

func asyncDouble(x int, ch chan int) { ch <- x * 2 }

func main() {
	ch := make(chan int)
	go asyncDouble(5, ch)
	result := <-ch
	fmt.Println(result) // 10
}
