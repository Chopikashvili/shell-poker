package game

import (
	"chopikashvili/shellpoker/card"
	"slices"
)

// Compares every player's best hand and determines the winner(s).
func DetermineWinners(hands []card.HandStrength) []int {
	slices.SortStableFunc(hands, card.CompareHands)
	slices.Reverse(hands)
	winners := []int{}
	//at least one hand must be the winning hand
	winners = append(winners, hands[0].PlayerId)
	for i := 1; i < len(hands); i++ {
		if card.CompareHands(hands[0], hands[i]) == 0 {
			winners = append(winners, hands[i].PlayerId)
		} else {
			break
		}
	}
	//standardizes the output
	slices.Sort(winners)
	return winners
}
