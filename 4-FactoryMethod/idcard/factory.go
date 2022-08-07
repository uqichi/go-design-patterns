package idcard

import "github.com/uqichi/go-design-patterns/4-FactoryMethod/framework"

type IDCardFactory struct {
	framework.Factory
	owners []string
}

func (fac IDCardFactory) CreateProduct(owner string) framework.Product {
	return NewIDCard(owner)
}
func (fac *IDCardFactory) RegisterProduct(product framework.Product) {
	fac.owners = append(fac.owners, product.(IDCard).owner)
}
