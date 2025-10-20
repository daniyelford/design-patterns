package main

import "fmt"

type Logger interface {
	Log(msg string)
}

type FileLogger struct{}

func (FileLogger) Log(msg string) { fmt.Println("[FILE]", msg) }

type NullLogger struct{}

func (NullLogger) Log(msg string) { /* do nothing */ }

type UserService struct{ logger Logger }

func (u UserService) Register(name string) {
	fmt.Println("Registering", name)
	u.logger.Log("User " + name + " registered")
}

func main() {
	s1 := UserService{logger: FileLogger{}}
	s1.Register("Ali")

	s2 := UserService{logger: NullLogger{}}
	s2.Register("Sara")
}
