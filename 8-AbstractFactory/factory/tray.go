package factory

type TrayIface interface {
	ItemIface
	Add(item ItemIface)
}

type Tray struct {
	Item
	values []ItemIface
}

func NewTray(item Item) Tray {
	return Tray{Item: item}
}

func (t *Tray) Add(item ItemIface) {
	t.values = append(t.values, item)
}

func (t *Tray) GetValues() []ItemIface {
	return t.values
}
