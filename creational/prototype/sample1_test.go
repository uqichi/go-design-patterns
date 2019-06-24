package prototype

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	cli := newClient()
	enemies := cli.createManyEnemies()
	for _, enemy := range enemies {
		fmt.Println(enemy.attack, enemy.defense)
	}
}
