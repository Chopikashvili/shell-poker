package card

import (
	"fmt"
	"math/rand"
)

type Card struct {
	value int  //Value of the card. From 2 to 15.
	suit  rune //Suit of the card. A unicode card suit character or a letter.
}

type Deck [52]Card

// Transforms the Card object into a readable string.
func (c Card) Read() string {
	var repr string
	//adding value
	highCards := [4]rune{'J', 'Q', 'K', 'A'}
	if c.value <= 10 {
		repr += fmt.Sprint(c.value)
	} else {
		repr += string(highCards[c.value-11])
	}
	//adding suit
	repr += string(c.suit)
	return repr
}

// Shuffles deck.
func (d Deck) Shuffle() {
	rand.Shuffle(52, func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func (d Deck) Deal(cardsUsed int) Card {
	card := d[cardsUsed]
	return card
}
