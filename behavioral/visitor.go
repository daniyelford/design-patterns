package main

import "fmt"

type Visitor interface {
	VisitCircle(Circle)
	VisitSquare(Square)
}

type Shape interface{ Accept(Visitor) }

type Circle struct{}

func (c Circle) Accept(v Visitor) { v.VisitCircle(c) }

type Square struct{}

func (s Square) Accept(v Visitor) { v.VisitSquare(s) }

type AreaVisitor struct{}

func (AreaVisitor) VisitCircle(c Circle) { fmt.Println("Circle area calculated") }
func (AreaVisitor) VisitSquare(s Square) { fmt.Println("Square area calculated") }

func main() {
	shapes := []Shape{Circle{}, Square{}}
	v := AreaVisitor{}
	for _, s := range shapes {
		s.Accept(v)
	}
}
