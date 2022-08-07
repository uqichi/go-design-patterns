package main

const (
	HandValueGuu = int(iota)
	HandValueCho
	HandValuePaa
)

var hand = []Hand{
	{HandValueGuu},
	{HandValueCho},
	{HandValuePaa},
}

var name = []string{
	"グー",
	"チョキ",
	"パー",
}

type Hand struct {
	handValue int
}

func GetHand(handValue int) Hand {
	return hand[handValue]
}

func (hand Hand) IsStrongerThan(h Hand) bool {
	return hand.fight(h) == 1
}

func (hand Hand) IsWeakerThan(h Hand) bool {
	return hand.fight(h) == -1
}

func (hand Hand) fight(h Hand) int {
	if hand == h {
		return 0
	} else if (hand.handValue+1)%3 == h.handValue {
		return 1
	}
	return -1
}

func (hand Hand) String() string {
	return name[hand.handValue]
}
