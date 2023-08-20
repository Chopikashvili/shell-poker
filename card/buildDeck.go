package card

import "errors"

func BuildDeck() (Deck, error) {
	d := Deck{}
	return buildDeck(d, false) //add support for letter suits later
}

func buildDeck(d Deck, letters bool) (Deck, error) {
	suits := [4]rune{'♠', '♣', '♥', '♦'}
	lettersuits := [4]rune{'S', 'C', 'H', 'D'}
	values := [13]int{}
	for i := 0; i < 13; i++ {
		values[i] = i + 2
	}
	for s := 0; s < 4; s++ {
		for v := 0; v < 13; v++ {
			if letters {
				d[s*13+v] = Card{suit: lettersuits[s], value: values[v]}
			} else {
				d[s*13+v] = Card{suit: suits[s], value: values[v]}
			}
		}
	}
	c := Card{suit: '♦', value: 14}
	if d[51] != c {
		return d, errors.New("Could not build deck")
	}
	return d, nil
}
