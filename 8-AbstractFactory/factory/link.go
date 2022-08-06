package factory

type LinkIface interface {
	ItemIface
}

type Link struct {
	Item
	url string
}

func NewLink(item Item, url string) Link {
	return Link{
		Item: item,
		url:  url,
	}
}

func (l Link) GetURL() string {
	return l.url
}
