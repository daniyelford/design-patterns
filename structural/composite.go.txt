package main

import "fmt"

// Component
type FileSystem interface {
	Show()
}

// Leaf
type File struct {
	name string
}

func (f File) Show() {
	fmt.Println("File:", f.name)
}

// Composite
type Folder struct {
	name     string
	children []FileSystem
}

func (f *Folder) Add(child FileSystem) {
	f.children = append(f.children, child)
}

func (f Folder) Show() {
	fmt.Println("Folder:", f.name)
	for _, child := range f.children {
		child.Show()
	}
}

func main() {
	file1 := File{"file1.txt"}
	file2 := File{"file2.txt"}

	folder1 := Folder{name: "Documents"}
	folder1.Add(file1)
	folder1.Add(file2)

	folder1.Show()
}
