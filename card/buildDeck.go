package card

import "errors"

//Constructs deck.
func BuildDeck(letters bool) (Deck, error) {
	d := Deck{}
	return buildDeck(d, letters) //add support for letter suits later
}

func buildDeck(d Deck, letters bool) (Deck, error) {
	//constructing lists of possible values and suits
	values := [13]int{}
	for i := 0; i < 13; i++ {
		values[i] = i + 2
	}
	//fills the deck
	for s := 0; s < 4; s++ {
		for v := 0; v < 13; v++ {
			if letters {
				d[s*13+v] = Card{suit: LetterSuits[s], value: values[v]}
			} else {
				d[s*13+v] = Card{suit: Suits[s], value: values[v]}
			}
		}
	}
	//checks if deck was constructed correctly
	c := Card{suit: '♦', value: 14}
	if d[51] != c {
		return d, errors.New("Could not build deck")
	}
	return d, nil
}
