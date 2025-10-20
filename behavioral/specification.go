package main

import "fmt"

type User struct {
	Age   int
	Email string
}

type Specification interface {
	IsSatisfiedBy(User) bool
}

type AgeAbove18 struct{}

func (AgeAbove18) IsSatisfiedBy(u User) bool { return u.Age > 18 }

type HasEmail struct{}

func (HasEmail) IsSatisfiedBy(u User) bool { return u.Email != "" }

type AndSpec struct{ a, b Specification }

func (s AndSpec) IsSatisfiedBy(u User) bool {
	return s.a.IsSatisfiedBy(u) && s.b.IsSatisfiedBy(u)
}

func main() {
	user := User{Age: 20, Email: "test@example.com"}
	spec := AndSpec{AgeAbove18{}, HasEmail{}}
	if spec.IsSatisfiedBy(user) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Rejected")
	}
}
