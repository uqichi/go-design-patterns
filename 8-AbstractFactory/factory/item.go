package factory

type ItemIface interface {
	MakeHTML() string
}

type Item struct {
	caption string
}

func NewItem(caption string) Item {
	return Item{caption}
}

func (i Item) GetCaption() string {
	return i.caption
}
