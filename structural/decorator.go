package main

import "fmt"

// Component
type Text interface {
	Render() string
}

// Concrete Component
type PlainText struct {
	text string
}

func (p PlainText) Render() string { return p.text }

// Decorator
type TextDecorator struct {
	wrapped Text
}

func (d TextDecorator) Render() string { return d.wrapped.Render() }

// Concrete Decorators
type BoldText struct {
	TextDecorator
}

func (b BoldText) Render() string {
	return "<b>" + b.wrapped.Render() + "</b>"
}

type ItalicText struct {
	TextDecorator
}

func (i ItalicText) Render() string {
	return "<i>" + i.wrapped.Render() + "</i>"
}

func main() {
	var text Text = PlainText{"Hello Decorator"}
	text = BoldText{TextDecorator{text}}
	text = ItalicText{TextDecorator{text}}

	fmt.Println(text.Render())
}
