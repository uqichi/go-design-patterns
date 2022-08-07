package main

import "fmt"

// Goは継承が使えないので派生型による共用と置換、多態性にインタフェースを使う
type AbstractDisplayIface interface {
	open()
	print()
	close()
	display(d AbstractDisplayIface)
}

// Goは継承が使えないのでコードの再利用にこれを埋め込む
type AbstractDisplay struct {
	AbstractDisplayIface
}

func (ad AbstractDisplay) display(d AbstractDisplayIface) {
	d.open()
	for i := 0; i < 5; i++ {
		d.print()
	}
	d.close()
}

type CharDisplay struct {
	AbstractDisplay
	ch rune
}

func newCharDisplay(ch rune) CharDisplay {
	return CharDisplay{
		ch: ch,
	}
}

func (cd CharDisplay) open() {
	fmt.Print("<<")
}

func (cd CharDisplay) print() {
	fmt.Print(cd.ch)
}

func (cd CharDisplay) close() {
	fmt.Println(">>")
}

type StringDisplay struct {
	AbstractDisplay
	str   string
	width int
}

func newStringDisplay(str string) StringDisplay {
	return StringDisplay{
		str:   str,
		width: len([]byte(str)),
	}
}

func (cd StringDisplay) open() {
	cd.printLine()
}

func (cd StringDisplay) print() {
	fmt.Printf("|%s|\n", cd.str)
}

func (cd StringDisplay) close() {
	cd.printLine()
}

func (cd StringDisplay) printLine() {
	for i := 0; i < cd.width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func main() {
	var d1 AbstractDisplayIface = newCharDisplay('H')
	var d2 AbstractDisplayIface = newStringDisplay("Hello, world")
	var d3 AbstractDisplayIface = newStringDisplay("こんにちは")
	d1.display(d1)
	d2.display(d2)
	d3.display(d3)
}
