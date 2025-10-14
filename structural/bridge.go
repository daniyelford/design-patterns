package main

import "fmt"

// Implementor
type Color interface {
	ApplyColor() string
}

// Concrete Implementor
type RedColor struct{}
func (RedColor) ApplyColor() string { return "Red" }

type BlueColor struct{}
func (BlueColor) ApplyColor() string { return "Blue" }

// Abstraction
type Shape struct {
	color Color
}

func (s *Shape) Draw() string { return "" } // abstract

// Refined Abstraction
type Circle struct{ Shape }
func (c Circle) Draw() string {
	return "Circle with color " + c.color.ApplyColor()
}

type Rectangle struct{ Shape }
func (r Rectangle) Draw() string {
	return "Rectangle with color " + r.color.ApplyColor()
}

func main() {
	redCircle := Circle{Shape{RedColor{}}}
	blueRect := Rectangle{Shape{BlueColor{}}}

	fmt.Println(redCircle.Draw())
	fmt.Println(blueRect.Draw())
}
