package main

import "fmt"

// Flyweight
type TreeType struct {
	Name  string
	Color string
}

func (t TreeType) Draw(x, y int) {
	fmt.Printf("Drawing %s (%s) at (%d, %d)\n", t.Name, t.Color, x, y)
}

// Factory
type TreeFactory struct {
	treeTypes map[string]*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{treeTypes: make(map[string]*TreeType)}
}

func (f *TreeFactory) GetTreeType(name, color string) *TreeType {
	key := name + color
	if _, ok := f.treeTypes[key]; !ok {
		f.treeTypes[key] = &TreeType{name, color}
	}
	return f.treeTypes[key]
}

func main() {
	factory := NewTreeFactory()

	tree1 := factory.GetTreeType("Oak", "Green")
	tree2 := factory.GetTreeType("Oak", "Green") // همان شیء قبلی

	tree1.Draw(10, 20)
	tree2.Draw(15, 25)
}
