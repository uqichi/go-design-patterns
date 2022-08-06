package listfactory

import "github.com/uqichi/go-design-patterns/8-AbstractFactory/factory"

type ListFactory struct{}

func (lf ListFactory) CreateLink(caption, url string) factory.LinkIface {
	return NewListLink(caption, url)
}

func (lf ListFactory) CreateTray(caption string) factory.TrayIface {
	return NewListTray(caption)
}

func (lf ListFactory) CreatePage(title, author string) factory.PageIface {
	return NewListPage(title, author)
}
