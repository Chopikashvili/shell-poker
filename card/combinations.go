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
	possibleHands := []HandStrength{}
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
						possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 8, CombName: "straight flush", OrderedCardValues: resultValues})
					} else {
						possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 5, CombName: "flush", OrderedCardValues: resultValues})
					}
				}
			}
			distribution := make([]int, 13)
			for i := 0; i < 13; i++ {
				if slices.Contains(resultValues, i+2) {
					distribution[i]++
				}
			}
			distributionCopy := distribution
			for i := 0; i < 4; i++ {
				distributionCopy = append(distributionCopy, distribution[i])
			}
			if containsSubslice(distributionCopy, []int{1, 1, 1, 1, 1}) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 4, CombName: "straight", OrderedCardValues: resultValues})
			} else if slices.Contains(distribution, 4) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 7, CombName: "four of a kind", OrderedCardValues: resultValues})
			} else if slices.Contains(distribution, 3) && slices.Contains(distribution, 2) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 6, CombName: "full house", OrderedCardValues: resultValues})
			} else if slices.Contains(distribution, 3) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 3, CombName: "three of a kind", OrderedCardValues: resultValues})
			} else if count(distribution, 2) == 2 {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 2, CombName: "two pairs", OrderedCardValues: resultValues})
			} else if slices.Contains(distribution, 2) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 1, CombName: "pair", OrderedCardValues: resultValues})
			} else {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 0, CombName: "high card", OrderedCardValues: resultValues})
			}
		}
	}
	return slices.MaxFunc(possibleHands, CompareHands), nil
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

func count(slice []int, comparable int) int {
	var counter int
	for _, elem := range slice {
		if elem == comparable {
			counter++
		}
	}
	return counter
}

func containsSubslice[T comparable](slice []T, subslice []T) bool {
	for i := 0; i < len(slice)-len(subslice); i++ {
		for j := i; i < len(subslice); j++ {
			if slices.Equal(slice[i:j+1], subslice) {
				return true
			}
		}
	}
	return false
}

// Sorts int arrays in reverse
func sortDesc(s []int) []int {
	slices.Sort(s)
	slices.Reverse(s)
	return s
}
