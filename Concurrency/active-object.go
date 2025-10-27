package main

import (
	"fmt"
	"time"
)

type Request struct {
	value int
	done  chan bool
}

type ActiveObject struct {
	queue chan Request
}

func (ao *ActiveObject) Start() {
	go func() {
		for req := range ao.queue {
			fmt.Println("Processing:", req.value)
			time.Sleep(time.Second)
			req.done <- true
		}
	}()
}

func (ao *ActiveObject) DoSomething(val int) {
	done := make(chan bool)
	ao.queue <- Request{val, done}
	<-done // optional: منتظر اتمام
}

func main() {
	ao := &ActiveObject{queue: make(chan Request, 10)}
	ao.Start()
	go ao.DoSomething(1)
	go ao.DoSomething(2)
	time.Sleep(3 * time.Second)
}
