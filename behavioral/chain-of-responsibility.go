package main

import "fmt"

type Handler interface {
	SetNext(Handler) Handler
	Handle(int)
}

type Base struct{ next Handler }

func (b *Base) SetNext(n Handler) Handler { b.next = n; return n }

type LowHandler struct{ Base }

func (h *LowHandler) Handle(a int) {
	if a <= 100 {
		fmt.Println("Low handled")
	}
	if h.next != nil {
		h.next.Handle(a)
	}
}

type HighHandler struct{ Base }

func (h *HighHandler) Handle(a int) {
	if a > 100 {
		fmt.Println("High handled")
	}
	if h.next != nil {
		h.next.Handle(a)
	}
}

func main() {
	h1 := &LowHandler{}
	h2 := &HighHandler{}
	h1.SetNext(h2)
	h1.Handle(150)
}
