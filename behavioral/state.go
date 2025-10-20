package main

import "fmt"

type State interface{ DoAction(*Context) }
type Context struct{ state State }

func (c *Context) SetState(s State) { c.state = s }

type StartState struct{}

func (StartState) DoAction(c *Context) { fmt.Println("Start"); c.SetState(StopState{}) }

type StopState struct{}

func (StopState) DoAction(c *Context) { fmt.Println("Stop") }

func main() {
	c := &Context{state: StartState{}}
	c.state.DoAction(c)
	c.state.DoAction(c)
}
