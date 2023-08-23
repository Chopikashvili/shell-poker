package card

import (
	"errors"
	"slices"
)

type HandStrength struct {
	PlayerId          int
	CombStrength      int
	CombName          string
	OrderedCardValues []int
}

func IdentifyCombinations(playerId int, hand []Card) (HandStrength, error) {
	if len(hand) != 7 {
		return HandStrength{playerId, -1, "unknown", []int{}}, errors.New("Something went wrong while calculating winner")
	}
	return identifyCombinations(playerId, hand)
}

func identifyCombinations(playerId int, hand []Card) (HandStrength, error) {
	suits := [4]rune{'♠', '♣', '♥', '♦'}
	for i := 0; i < 7; i++ {
		middle := slices.Delete(hand, i, i+1)
		for j := 0; j < 6; j++ {
			resultHand := slices.Delete(middle, j, j+1)
			resultValues := []int{}
			for _, c := range resultHand {
				resultValues = append(resultValues, c.value)
				resultValues = sortDesc(resultValues)
			}
			for _, r := range suits {
				if len(filter(resultHand, func(c Card) bool { return c.suit == r })) == 5 {
					if resultValues[0]-resultValues[4] == 4 {
						return HandStrength{PlayerId: playerId, CombStrength: 8, CombName: "straight flush", OrderedCardValues: resultValues}, nil
					} else {
						return HandStrength{PlayerId: playerId, CombStrength: 5, CombName: "flush", OrderedCardValues: resultValues}, nil
					}
				}
			}
		}
	}
	return HandStrength{}, nil //temp
}

func filter[T any](slice []T, predicate func(T) bool) []T {
	subslice := []T{}
	for _, elem := range slice {
		if predicate(elem) {
			subslice = append(subslice, elem)
		}
	}
	return subslice
}

// Sorts int arrays in reverse
func sortDesc(s []int) []int {
	slices.Sort(s)
	slices.Reverse(s)
	return s
}
