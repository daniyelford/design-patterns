package main

import "fmt"

// کلاس قدیمی
type OldPrinter struct{}

func (o OldPrinter) PrintText(text string) {
	fmt.Println("Printing:", text)
}

// رابط مورد انتظار
type Printer interface {
	Print(text string)
}

// Adapter
type PrinterAdapter struct {
	oldPrinter OldPrinter
}

func (p PrinterAdapter) Print(text string) {
	p.oldPrinter.PrintText(text)
}

func main() {
	old := OldPrinter{}
	adapter := PrinterAdapter{oldPrinter: old}
	adapter.Print("Hello Adapter!")
}
