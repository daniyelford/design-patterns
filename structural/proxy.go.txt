package main

import "fmt"

// Subject
type Image interface {
	Display()
}

// RealSubject
type RealImage struct {
	filename string
}

func NewRealImage(filename string) *RealImage {
	img := &RealImage{filename}
	img.loadFromDisk()
	return img
}

func (r *RealImage) loadFromDisk() {
	fmt.Println("Loading", r.filename)
}

func (r *RealImage) Display() {
	fmt.Println("Displaying", r.filename)
}

// Proxy
type ProxyImage struct {
	filename  string
	realImage *RealImage
}

func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename: filename}
}

func (p *ProxyImage) Display() {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.filename)
	}
	p.realImage.Display()
}

func main() {
	image := NewProxyImage("photo.jpg")

	image.Display()
	image.Display()
}
