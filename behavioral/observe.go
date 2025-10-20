package main

import "fmt"

type Observer interface{ Update(event string, data any) }
type EmailNotifier struct{}

func (EmailNotifier) Update(e string, d any) { fmt.Println("Email:", e) }

type Logger struct{}

func (Logger) Update(e string, d any) { fmt.Println("Log:", e) }

type Subject struct{ observers []Observer }

func (s *Subject) Attach(o Observer) { s.observers = append(s.observers, o) }
func (s *Subject) Notify(e string, d any) {
	for _, o := range s.observers {
		o.Update(e, d)
	}
}

func Main() {
	s := &Subject{}
	s.Attach(EmailNotifier{})
	s.Attach(Logger{})
	s.Notify("order.created", map[string]int{"id": 123})
}
