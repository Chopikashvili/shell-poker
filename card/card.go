package card

import (
	"fmt"
	"math/rand"
)

type Card struct {
	value int //From 2 to 15.
	suit  rune
}

type Deck [52]Card

func (c Card) Read() string {
	var repr string
	highCards := [4]rune{'J', 'Q', 'K', 'A'}
	if c.value <= 10 {
		repr += fmt.Sprint(c.value)
	} else {
		repr += string(highCards[c.value-11])
	}
	repr += string(c.suit)
	return repr
}

func (d Deck) Shuffle() {
	rand.Shuffle(52, func(i, j int) { d[i], d[j] = d[j], d[i] })
}
