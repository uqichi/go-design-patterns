package listfactory

import (
	"fmt"

	"github.com/uqichi/go-design-patterns/8-AbstractFactory/factory"
)

type ListLink struct {
	*factory.Link
}

func NewListLink(caption, url string) ListLink {
	list := factory.NewLink(factory.NewItem(caption), url)
	return ListLink{
		Link: &list,
	}
}

func (ll ListLink) MakeHTML() string {
	return fmt.Sprintf("  <li><a href=\"%s\">%s</a></li>\n", ll.GetURL(), ll.GetCaption())
}
