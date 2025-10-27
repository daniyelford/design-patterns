package main

import (
	"fmt"
	"time"
)

var ready = false
var signal = make(chan bool)

func worker() {
	for !ready {
		fmt.Println("waiting...")
		<-signal // منتظر می‌ماند تا ready شود
	}
	fmt.Println("now working!")
}

func main() {
	go worker()
	time.Sleep(2 * time.Second)
	ready = true
	signal <- true // اجازه اجرا
	time.Sleep(time.Second)
}
