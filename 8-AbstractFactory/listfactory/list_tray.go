package listfactory

import (
	"fmt"

	"github.com/uqichi/go-design-patterns/8-AbstractFactory/factory"
)

type ListTray struct {
	*factory.Tray
}

func NewListTray(caption string) ListTray {
	tray := factory.NewTray(factory.NewItem(caption))
	return ListTray{
		Tray: &tray,
	}
}

func (lt ListTray) MakeHTML() string {
	buf := "<li>\n"
	buf += fmt.Sprintf("%s\n", lt.GetCaption())
	buf += "<ul>\n"
	for _, value := range lt.GetValues() {
		buf += value.MakeHTML()
	}
	buf += "</ul>\n"
	buf += "</li>\n"
	return buf
}
