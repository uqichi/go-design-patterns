package prototype

import "fmt"

type cloneable interface {
	clone() cloneable
}

type enemy struct {
	attack  int
	defense int
}

func newEnemy(attack, defense int) *enemy {
	return &enemy{
		attack:  attack,
		defense: defense,
	}
}

func (e *enemy) clone() cloneable {
	ne := newEnemy(e.attack, e.defense)
	return ne
}

type client struct{}

func newClient() *client {
	return new(client)
}

func (cli *client) createManyEnemies() []*enemy {
	prototype := newEnemy(50, 50)
	cli.configure(prototype)

	enemies := make([]*enemy, 10)

	for i := range enemies {
		enemies[i] = prototype.clone().(*enemy)
	}

	return enemies
}

func (cli *client) configure(enemy *enemy) {
	fmt.Println("process very very complicated high-cost configuration")
}
