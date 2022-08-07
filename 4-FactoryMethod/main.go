package main

import (
	"github.com/uqichi/go-design-patterns/4-FactoryMethod/framework"
	"github.com/uqichi/go-design-patterns/4-FactoryMethod/idcard"
)

func main() {
	var fac framework.FactoryIface = &idcard.IDCardFactory{}
	var card1 framework.Product = fac.GetProduct(fac, "John")
	var card2 framework.Product = fac.GetProduct(fac, "Michel")
	var card3 framework.Product = fac.GetProduct(fac, "Nick")
	card1.Use()
	card2.Use()
	card3.Use()
}
