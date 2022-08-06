package listfactory

import (
	"fmt"

	"github.com/uqichi/go-design-patterns/8-AbstractFactory/factory"
)

type ListPage struct {
	*factory.Page
}

func NewListPage(title, author string) ListPage {
	page := factory.NewPage(title, author)
	return ListPage{
		Page: &page,
	}
}

func (lp ListPage) MakeHTML() string {
	buf := "<html>\n"
	buf += fmt.Sprintf("  <head><title>%s</title></head>\n", lp.GetTitle())
	buf += "<body>\n"
	buf += fmt.Sprintf("<h1>%s</h1>", lp.GetTitle())
	buf += "<ul>"
	for _, item := range lp.GetContent() {
		buf += item.MakeHTML()
	}
	buf += "</ul>"
	buf += fmt.Sprintf("<hr><adress>%s</adress>", lp.GetAuthor())
	buf += "</body>\n</html>\n"
	return buf
}
