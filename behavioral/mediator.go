package main

import "fmt"

type Mediator interface {
	Notify(sender *User, msg string)
}

type ChatRoom struct {
	users []*User
}

func (c *ChatRoom) Register(u *User) {
	c.users = append(c.users, u)
	u.mediator = c
}
func (c *ChatRoom) Notify(sender *User, msg string) {
	for _, u := range c.users {
		if u != sender {
			u.Receive(msg)
		}
	}
}

type User struct {
	name     string
	mediator Mediator
}

func (u *User) Send(msg string) {
	fmt.Println(u.name, "sends:", msg)
	u.mediator.Notify(u, msg)
}
func (u *User) Receive(msg string) {
	fmt.Println(u.name, "got:", msg)
}

func main() {
	chat := &ChatRoom{}
	ali := &User{name: "Ali"}
	sara := &User{name: "Sara"}
	chat.Register(ali)
	chat.Register(sara)
	ali.Send("Salam!")
}
