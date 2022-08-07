package main

import (
	"errors"

	"github.com/uqichi/go-design-patterns/8-AbstractFactory/factory"
	"github.com/uqichi/go-design-patterns/8-AbstractFactory/listfactory"
)

func main() {
	fac, err := getFactory("list")
	if err != nil {
		panic(err)
	}
	asahi := fac.CreateLink("Asahi", "http://www.asahi.com")
	yomiuri := fac.CreateLink("Yomiuri", "http://www.yomiuri.co.jp")
	usYahoo := fac.CreateLink("Yahoo", "http://www.yahoo.com")
	jaYahoo := fac.CreateLink("Yahoo!Japan", "http://www.yahoo.co.jp")
	google := fac.CreateLink("Google", "http://www.google.com")
	excite := fac.CreateLink("Excite", "http://www.excite.co.jp")

	traynews := fac.CreateTray("Newspaper")
	traynews.Add(asahi)
	traynews.Add(yomiuri)

	trayyahoo := fac.CreateTray("Yahoo!")
	trayyahoo.Add(usYahoo)
	trayyahoo.Add(jaYahoo)

	traysearch := fac.CreateTray("Search Engine")
	traysearch.Add(trayyahoo)
	traysearch.Add(excite)
	traysearch.Add(google)

	page := fac.CreatePage("LinkPage", "Yukichi")
	page.Add(traynews)
	page.Add(traysearch)
	page.Output(page)
}

func getFactory(name string) (factory.Factory, error) {
	switch name {
	case "list":
		return listfactory.ListFactory{}, nil
	}
	return nil, errors.New("factory not found")
}
