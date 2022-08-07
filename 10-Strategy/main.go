package main

import "fmt"

func main() {
	player1 := NewPlayer("馬鹿", NewWinningStrategy(314))
	player2 := NewPlayer("天才", NewProbStrategy(15))
	for i := 0; i < 10000; i++ {
		nextHand1 := player1.NextHand()
		nextHand2 := player2.NextHand()
		if nextHand1.IsStrongerThan(nextHand2) {
			fmt.Println("Winner", player1)
			player1.Win()
			player2.Lose()
		} else if nextHand2.IsStrongerThan(nextHand1) {
			fmt.Println("Winner", player2)
			player1.Lose()
			player2.Win()
		} else {
			fmt.Println("Even...")
			player1.Even()
			player2.Even()
		}
		fmt.Println("Total result:")
		fmt.Println(player1)
		fmt.Println(player2)
	}
}
