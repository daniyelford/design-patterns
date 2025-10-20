package main

import "fmt"

type Expression interface{ Interpret() int }

type Number struct{ value int }

func (n Number) Interpret() int { return n.value }

type Plus struct{ left, right Expression }

func (p Plus) Interpret() int { return p.left.Interpret() + p.right.Interpret() }

func main() {
	expr := Plus{Number{5}, Number{10}}
	fmt.Println(expr.Interpret()) // 15
}
