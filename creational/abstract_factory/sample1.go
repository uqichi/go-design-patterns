package abstract_factory

import "fmt"

type protein interface {
	getProteinName() string
}

type chicken struct{}

func (*chicken) getProteinName() string {
	return "chicken"
}

type beef struct{}

func (*beef) getProteinName() string {
	return "beef"
}

type pork struct{}

func (*pork) getProteinName() string {
	return "pork"
}

type vegetable interface {
	getVegetableName() string
}

type eggPlant struct{}

func (*eggPlant) getVegetableName() string {
	return "egg plant"
}

type bean struct{}

func (*bean) getVegetableName() string {
	return "bean"
}

type asparagus struct{}

func (*asparagus) getVegetableName() string {
	return "asparagus"
}

type curry struct {
	protein    protein
	vegetables []vegetable
}

func newCurry(protein protein, vegetables []vegetable) *curry {
	return &curry{
		protein:    protein,
		vegetables: vegetables,
	}
}

func (curry *curry) taste() {
	fmt.Println("taste like " + curry.protein.getProteinName())
	vegeTaste := ""
	for _, vege := range curry.vegetables {
		vegeTaste += vege.getVegetableName() + " "
	}
	fmt.Println("taste like " + vegeTaste)
}

type curryFactory interface {
	getProtein() protein
	getVegetables() []vegetable
}

type vegeChickenCurryFactory struct{}

func (*vegeChickenCurryFactory) getProtein() protein {
	return new(chicken)
}

func (*vegeChickenCurryFactory) getVegetables() []vegetable {
	return []vegetable{
		new(eggPlant),
		new(bean),
		new(asparagus),
	}
}

type keemaCurryFactory struct{}

func (*keemaCurryFactory) getProtein() protein {
	return new(beef)
}

func (*keemaCurryFactory) getVegetables() []vegetable {
	return []vegetable{
		new(bean),
		new(eggPlant),
	}
}

type porkAsparagusCurryFactory struct{}

func (*porkAsparagusCurryFactory) getProtein() protein {
	return new(pork)
}

func (*porkAsparagusCurryFactory) getVegetables() []vegetable {
	return []vegetable{
		new(asparagus),
	}
}

func createCurryFactory(name string) curryFactory {
	switch name {
	case "vege_chicken":
		return &vegeChickenCurryFactory{}
	case "keema":
		return &keemaCurryFactory{}
	case "pork_asparagus":
		return &porkAsparagusCurryFactory{}
	default:
		return nil
	}
}
