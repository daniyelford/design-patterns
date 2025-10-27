package main

import (
	"fmt"
	"time"
)

type EventHandler func()

func reactor(events map[string]EventHandler) {
	for {
		for name, handler := range events {
			fmt.Printf("⏺ Triggering: %s\n", name)
			handler()
			time.Sleep(time.Second)
		}
	}
}

func main() {
	events := map[string]EventHandler{
		"read":  func() { fmt.Println("📥 reading data") },
		"write": func() { fmt.Println("📤 writing data") },
	}
	reactor(events)
}
