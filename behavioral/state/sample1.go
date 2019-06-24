package state

import "fmt"

type state interface {
	getWear() string
}

type rainyState struct{}

func newRainyState() *rainyState {
	return &rainyState{}
}

func (*rainyState) getWear() string {
	return "waterproof coat"
}

type sunnyState struct{}

func newSunnyState() *sunnyState {
	return &sunnyState{}
}

func (*sunnyState) getWear() string {
	return "simple t-shirt"
}

type statePattern struct {
	state state
}

func newStatePattern() *statePattern {
	return &statePattern{}
}

func (sp *statePattern) setState(state state) {
	sp.state = state
}

func (sp *statePattern) whatShouldIWear() {
	fmt.Println("I should wear " + sp.state.getWear())
}
