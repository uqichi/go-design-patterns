package main

import (
	"math/rand"
)

type WinningStrategy struct {
	random   *rand.Rand
	won      bool
	prevHand Hand
}

func NewWinningStrategy(seed int64) WinningStrategy {
	return WinningStrategy{
		random: rand.New(rand.NewSource(seed)),
	}
}

func (w WinningStrategy) NextHand() Hand {
	if !w.won {
		w.prevHand = GetHand(w.random.Intn(3))
	}
	return w.prevHand
}

func (w WinningStrategy) Study(win bool) {
	w.won = win
}
