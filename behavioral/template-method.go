package main

import "fmt"

type Game interface{ Play() }
type BaseGame struct{}

func (BaseGame) Play() {}

// In Go معمولاً از الگوی دیگری استفاده می‌کنیم؛ با ترکیب (composition) و فراخوانی توالی توابع:
type Chess struct{}

func (Chess) start()    { fmt.Println("Chess starts") }
func (Chess) playTurn() { fmt.Println("Move") }
func (c Chess) Play()   { c.start(); c.playTurn(); fmt.Println("Game over") }

type Football struct{}

func (Football) start()    { fmt.Println("Kick off!") }
func (Football) playTurn() { fmt.Println("Pass ball!") }
func (Football) end()      { fmt.Println("Game over") }
func (f Football) Play() {
	f.start()
	f.playTurn()
	f.end()
}

// func main() { Chess{}.Play() }
func main() {
	var games []Game = []Game{Chess{}, Football{}}
	for _, g := range games {
		g.Play()
		fmt.Println("---")
	}
}
