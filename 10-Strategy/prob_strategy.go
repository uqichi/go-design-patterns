package main

import "math/rand"

type ProbStrategy struct {
	random           *rand.Rand
	prevHandValue    int
	currentHandValue int
	history          [][]int
}

func NewProbStrategy(seed int64) ProbStrategy {
	return ProbStrategy{
		random: rand.New(rand.NewSource(seed)),
		// [前回に出した手][今回出す手]
		history: [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
	}
}

func (p ProbStrategy) NextHand() Hand {
	bet := p.random.Intn(p.getSum(p.currentHandValue))
	handValue := 0
	if bet < p.history[p.currentHandValue][0] {
		handValue = 0
	} else if bet < p.history[p.currentHandValue][0]+p.history[p.currentHandValue][1] {
		handValue = 1
	} else {
		handValue = 2
	}
	p.prevHandValue = p.currentHandValue
	p.currentHandValue = handValue
	return GetHand(handValue)
}

func (p ProbStrategy) getSum(hv int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += p.history[hv][i]
	}
	return sum
}

func (p ProbStrategy) Study(win bool) {
	if win {
		p.history[p.prevHandValue][p.currentHandValue]++
	} else {
		p.history[p.prevHandValue][(p.currentHandValue+1)%3]++
		p.history[p.prevHandValue][(p.currentHandValue+2)%3]++
	}
}
