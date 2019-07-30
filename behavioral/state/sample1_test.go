package state

import "testing"

func Test(t *testing.T) {
	sp := newStatePattern()

	rainy := newRainyState()
	sp.setState(rainy)
	sp.whatShouldIWear()

	sunny := newSunnyState()
	sp.setState(sunny)
	sp.whatShouldIWear()
}
