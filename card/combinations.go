package card

import (
	"chopikashvili/shellpoker/general"
	"errors"
	"slices"
)

type HandStrength struct {
	PlayerId          int
	CombStrength      int
	CombName          string
	OrderedCardValues []int
}

// Checks for combinations in a hand
func IdentifyCombinations(playerId int, hand []Card) (HandStrength, error) {
	if len(hand) != 7 {
		return HandStrength{playerId, -1, "unknown", []int{}}, errors.New("Something went wrong while calculating winner")
	}
	return identifyCombinations(playerId, hand)
}

func identifyCombinations(playerId int, hand []Card) (HandStrength, error) {
	possibleHands := []HandStrength{}
	suits := [4]rune{'♠', '♣', '♥', '♦'}
	for i := 0; i < 6; i++ {
		for j := i + 1; j < 7; j++ {
			//makes a hand
			resultHand := make([]Card, 0)
			for n, c := range hand {
				if n != i && n != j {
					resultHand = append(resultHand, c)
				}
			}
			//makes a slice of just the values in the hand, plus a copy of a hand, and sorts the original by descending order
			resultValues := make([]int, 0)
			for _, c := range resultHand {
				resultValues = append(resultValues, c.value)
			}
			resultValuesDesc := make([]int, 5)
			resultValues, resultValuesDesc = general.SortDesc(resultValues)
			//checks for flush and straight flush
			for _, r := range suits {
				if len(general.Filter(resultHand, func(c Card) bool { return c.suit == r })) == 5 {
					if resultValuesDesc[0]-resultValuesDesc[4] == 4 {
						possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 8, CombName: "straight flush", OrderedCardValues: resultValuesDesc})
					} else {
						possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 5, CombName: "flush", OrderedCardValues: resultValuesDesc})
					}
				}
			}
			//makes distribution of values, plus a copy for determining straights with a K-A-2 sequence
			distribution := make([]int, 13)
			distributionCopy := make([]int, 17)
			for i := 0; i < 13; i++ {
				if slices.Contains(resultValues, i+2) {
					distribution[i]++
					distributionCopy[i]++
				}
			}
			for i := 0; i < 4; i++ {
				distributionCopy[i+13] = distributionCopy[i]
			}
			//checks for all non-flush values. Optimization is welcome
			if general.ContainsSubslice(distributionCopy, []int{1, 1, 1, 1, 1}) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 4, CombName: "straight", OrderedCardValues: resultValuesDesc})
			} else if slices.Contains(distribution, 4) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 7, CombName: "four of a kind", OrderedCardValues: resultValuesDesc})
			} else if slices.Contains(distribution, 3) && slices.Contains(distribution, 2) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 6, CombName: "full house", OrderedCardValues: resultValuesDesc})
			} else if slices.Contains(distribution, 3) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 3, CombName: "three of a kind", OrderedCardValues: resultValuesDesc})
			} else if general.Count(distribution, 2) == 2 {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 2, CombName: "two pairs", OrderedCardValues: resultValuesDesc})
			} else if slices.Contains(distribution, 2) {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 1, CombName: "pair", OrderedCardValues: resultValuesDesc})
			} else {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 0, CombName: "high card", OrderedCardValues: resultValuesDesc})
			}
		}
	}
	//finds the most valuable possible hand and returns it
	return slices.MaxFunc(possibleHands, CompareHands), nil
}
