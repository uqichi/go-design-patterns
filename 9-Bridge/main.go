package main

import "fmt"

// Abstraction
// 機能のクラス階層の最上位
type Display struct {
	impl DisplayImpl
}

func NewDisplay(impl DisplayImpl) Display {
	return Display{impl}
}

func (d Display) Open() {
	d.impl.RawOpen()
}

func (d Display) Print() {
	d.impl.RawPrint()
}

func (d Display) Close() {
	d.impl.RawClose()
}

func (d Display) DoDisplay() {
	d.Open()
	d.Print()
	d.Close()
}

// RefinedAbstraction
// 機能のクラス階層 Displayの拡張機能
type CountDisplay struct {
	Display
}

func NewCountDisplay(impl DisplayImpl) CountDisplay {
	return CountDisplay{NewDisplay(impl)}
}

func (cd CountDisplay) DoMultiDisplay(times int) {
	cd.Open()
	for i := 0; i < times; i++ {
		cd.Print()
	}
	cd.Close()
}

// Implementor
// 実装のクラス階層の最上位
type DisplayImpl interface {
	RawOpen()
	RawPrint()
	RawClose()
}

// ConcreateImplementor
// 実装のクラス階層 具体的な実装
type StringDisplayImpl struct {
	str string
}

func NewStringDisplayImpl(str string) StringDisplayImpl {
	return StringDisplayImpl{str}
}

func (sdi StringDisplayImpl) RawOpen() {
	fmt.Println(">----------")
}

func (sdi StringDisplayImpl) RawPrint() {
	fmt.Printf("|%s|\n", sdi.str)
}

func (sdi StringDisplayImpl) RawClose() {
	fmt.Println("----------<")
}

func main() {
	var d1 Display = NewDisplay(NewStringDisplayImpl("Hello, Japan"))
	var d2 Display = NewCountDisplay(NewStringDisplayImpl("Hello, World")).Display
	var d3 CountDisplay = NewCountDisplay(NewStringDisplayImpl("Hello, Universe"))
	d1.DoDisplay()
	d2.DoDisplay()
	d3.DoDisplay()
	d3.DoMultiDisplay(5)
}
