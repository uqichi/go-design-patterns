package abstract_factory

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	name := "vege_chicken"
	fmt.Println(name)
	factory := createCurryFactory(name)
	curry := newCurry(factory.getProtein(), factory.getVegetables())
	curry.taste()
}

func Test2(t *testing.T) {
	name := "keema"
	fmt.Println(name)
	factory := createCurryFactory(name)
	curry := newCurry(factory.getProtein(), factory.getVegetables())
	curry.taste()
}

func Test3(t *testing.T) {
	name := "pork_asparagus"
	fmt.Println(name)
	factory := createCurryFactory(name)
	curry := newCurry(factory.getProtein(), factory.getVegetables())
	curry.taste()
}
