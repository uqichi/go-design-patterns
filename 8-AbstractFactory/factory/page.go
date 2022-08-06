package factory

import (
	"fmt"
	"os"
)

type PageIface interface {
	Add(item ItemIface)
	Output(pi PageIface)
	MakeHTML() string
}

type Page struct {
	title   string
	author  string
	content []ItemIface
}

func NewPage(title, author string) Page {
	return Page{title: title, author: author}
}

func (p *Page) Add(item ItemIface) {
	p.content = append(p.content, item)
}

func (p Page) Output(pi PageIface) {
	filename := p.title + ".html"
	file, _ := os.Create(filename)
	defer file.Close()
	b := []byte(pi.MakeHTML())
	_, _ = file.Write(b)
	fmt.Printf("%sを作成しました\n", filename)
}

func (p Page) GetTitle() string {
	return p.title
}

func (p Page) GetAuthor() string {
	return p.author
}

func (p Page) GetContent() []ItemIface {
	return p.content
}
