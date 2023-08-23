package game

import (
	"chopikashvili/shellpoker/card"
	"slices"
)

func DetermineWinners(hands []card.HandStrength) []int {
	slices.SortStableFunc(hands, card.CompareHands)
	slices.Reverse(hands)
	winners := []int{}
	winners = append(winners, hands[0].PlayerId)
	for i := 1; i < len(hands); i++ {
		if card.CompareHands(hands[0], hands[i]) == 0 {
			winners = append(winners, hands[i].PlayerId)
		} else {
			break
		}
	}
	return winners
}
