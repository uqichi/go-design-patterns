package framework

type FactoryIface interface {
	GetProduct(iface FactoryIface, owner string) Product
	CreateProduct(owner string) Product
	RegisterProduct(product Product)
}

type Factory struct {
	FactoryIface
}

func (fac Factory) GetProduct(iface FactoryIface, owner string) Product {
	p := iface.CreateProduct(owner)
	iface.RegisterProduct(p)
	return p
}
