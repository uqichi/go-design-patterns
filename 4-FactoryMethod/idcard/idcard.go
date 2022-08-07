package idcard

import (
	"fmt"

	"github.com/uqichi/go-design-patterns/4-FactoryMethod/framework"
)

type IDCard struct {
	framework.Product
	owner string
}

func NewIDCard(owner string) IDCard {
	fmt.Println(owner, "のカードを作ります")
	return IDCard{owner: owner}
}

func (card IDCard) Use() {
	fmt.Println(card.owner, "のカードを使います")
}
