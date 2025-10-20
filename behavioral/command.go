package main

import "fmt"

type Command interface{ Execute() }
type Light struct{}

func (Light) On()  { fmt.Println("Light on") }
func (Light) Off() { fmt.Println("Light off") }

type LightOnCommand struct{ light Light }

func (c LightOnCommand) Execute() { c.light.On() }

type Remote struct{}

func (Remote) Submit(cmd Command) { cmd.Execute() }

func main() {
	light := Light{}
	cmd := LightOnCommand{light}
	Remote{}.Submit(cmd)
}
