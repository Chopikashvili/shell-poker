package game

import (
	"chopikashvili/shellpoker/card"
	"slices"
)

func DetermineWinners(hands []card.HandStrength) []int {
	slices.SortStableFunc(hands, CompareHands)
	winners := []int{}
	winners = append(winners, hands[0].PlayerId)
	for i := 1; i < len(hands); i++ {
		if CompareHands(hands[0], hands[i]) == 0 {
			winners = append(winners, hands[i].PlayerId)
		} else {
			break
		}
	}
	return winners
}

func CompareHands(a, b card.HandStrength) int {
	if a.CombStrength == b.CombStrength {
		for i := 0; i < 5; i++ {
			if a.OrderedCardValues[i]-b.OrderedCardValues[i] != 0 {
				return b.OrderedCardValues[i] - a.OrderedCardValues[i]
			}
		}
		return 0
	}
	return b.CombStrength - a.CombStrength
}
