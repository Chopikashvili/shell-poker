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

type Combination struct {
	strength  int
	name      string
	condition func([]Card, []int, []int, []int) bool
}

var combinations = []Combination{
	{strength: 8, name: "straight flush", condition: func(cards []Card, i, dis, disC []int) bool {
		for _, r := range CombinedSuits {
			if len(general.Filter(cards, func(c Card) bool { return c.suit == r })) == 5 && i[0]-i[4] == 4 {
				return true
			}
		}
		return false
	}},
	{strength: 7, name: "four of a kind", condition: func(cards []Card, i, dis, disC []int) bool { return slices.Contains(dis, 4) }},
	{strength: 6, name: "full house", condition: func(cards []Card, i, dis, disC []int) bool { return slices.Contains(dis, 3) && slices.Contains(dis, 2) }},
	{strength: 5, name: "flush", condition: func(cards []Card, i, dis, disC []int) bool {
		for _, r := range CombinedSuits {
			if len(general.Filter(cards, func(c Card) bool { return c.suit == r })) == 5 {
				return true
			}
		}
		return false
	}},
	{strength: 4, name: "straight", condition: func(cards []Card, i, dis, disC []int) bool {
		return general.ContainsSubslice(disC, []int{1, 1, 1, 1, 1})
	}},
	{strength: 3, name: "three of a kind", condition: func(cards []Card, i, dis, disC []int) bool { return slices.Contains(dis, 3) }},
	{strength: 2, name: "two pairs", condition: func(cards []Card, i, dis, disC []int) bool { return general.Count(dis, 2) == 2 }},
	{strength: 1, name: "pair", condition: func(cards []Card, i, dis, disC []int) bool { return slices.Contains(dis, 2) }},
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
			flag := false
			for _, c := range combinations {
				if c.condition(resultHand, resultValuesDesc, distribution, distributionCopy) {
					possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: c.strength, CombName: c.name, OrderedCardValues: resultValuesDesc})
					flag = true
					break
				}
			}
			if !flag {
				possibleHands = append(possibleHands, HandStrength{PlayerId: playerId, CombStrength: 0, CombName: "high card", OrderedCardValues: resultValuesDesc})
			}
		}
	}
	//finds the most valuable possible hand and returns it
	return slices.MaxFunc(possibleHands, CompareHands), nil
}
