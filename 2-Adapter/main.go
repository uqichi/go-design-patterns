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
// Bannerがコードから隠されていることが重要
func main() {
	var p Print = PrintBanner{"Hello"}
	p.printWeak()
	p.printStrong()
}
