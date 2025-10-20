package main

import "fmt"

type Memento struct{ text string }

func (m Memento) GetText() string { return m.text }

type Editor struct{ text string }

func (e *Editor) Type(t string)     { e.text += " " + t }
func (e *Editor) GetText() string   { return e.text }
func (e *Editor) Save() Memento     { return Memento{text: e.text} }
func (e *Editor) Restore(m Memento) { e.text = m.GetText() }

func main() {
	e := &Editor{}
	e.Type("Salam")
	saved := e.Save()
	e.Type("chetori?")
	fmt.Println(e.GetText())
	e.Restore(saved)
	fmt.Println(e.GetText())
}
