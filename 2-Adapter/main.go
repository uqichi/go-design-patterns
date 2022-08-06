package main

import (
	"fmt"
)

// Adaptee
type Banner string

func (b Banner) showWithParen() {
	fmt.Printf("(%s)\n", b)
}
func (b Banner) showWithAster() {
	fmt.Printf("*%s*\n", b)
}

type Print interface {
	printWeak()
	printStrong()
}

// Adapter
type PrintBanner struct {
	Banner
}

func (pb PrintBanner) printWeak() {
	pb.Banner.showWithParen()
}

func (pb PrintBanner) printStrong() {
	pb.Banner.showWithAster()
}

// Client
// it is import that Banner is hidden from the code
func main() {
	var p Print = PrintBanner{"Hello"}
	p.printWeak()
	p.printStrong()
}

// Output
//(Hello)
//*Hello*
