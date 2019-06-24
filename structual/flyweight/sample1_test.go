package flyweight

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	factory := GetFactoryInstance()

	fmt.Println(factory.GetLetter("A").Lower)
	fmt.Println(factory.GetLetter("G").Upper)
}
