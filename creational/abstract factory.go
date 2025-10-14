package main

import "fmt"

// Button interface
type Button interface {
	Render()
}

// Checkbox interface
type Checkbox interface {
	Render()
}

// Windows implementations
type WindowsButton struct{}
func (WindowsButton) Render() { fmt.Println("Render Windows Button") }

type WindowsCheckbox struct{}
func (WindowsCheckbox) Render() { fmt.Println("Render Windows Checkbox") }

// Mac implementations
type MacButton struct{}
func (MacButton) Render() { fmt.Println("Render Mac Button") }

type MacCheckbox struct{}
func (MacCheckbox) Render() { fmt.Println("Render Mac Checkbox") }

// Abstract Factory
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Windows Factory
type WindowsFactory struct{}
func (WindowsFactory) CreateButton() Button { return WindowsButton{} }
func (WindowsFactory) CreateCheckbox() Checkbox { return WindowsCheckbox{} }

// Mac Factory
type MacFactory struct{}
func (MacFactory) CreateButton() Button { return MacButton{} }
func (MacFactory) CreateCheckbox() Checkbox { return MacCheckbox{} }

// client
func RenderUI(factory GUIFactory) {
	btn := factory.CreateButton()
	chk := factory.CreateCheckbox()
	btn.Render()
	chk.Render()
}

func main() {
	os := "mac"

	var factory GUIFactory
	if os == "windows" {
		factory = WindowsFactory{}
	} else {
		factory = MacFactory{}
	}

	RenderUI(factory)
}
