package factory

type Factory interface {
	CreateLink(caption, url string) LinkIface
	CreateTray(caption string) TrayIface
	CreatePage(title, author string) PageIface
}
