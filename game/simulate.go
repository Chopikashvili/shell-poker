package game

import (
	"chopikashvili/shellpoker/card"
	"slices"
)

func simulateGame(robot Player, deck card.Deck, playerNumber int, community []card.Card) (bool, error) {
	cardsUsed := 0
	deck.Shuffle()
	simCommunity := slices.Clone(community)
	for i := len(simCommunity); i < 5; i++ {
		simCommunity = append(simCommunity, simulateDeal(deck, &cardsUsed, robot.Hand))
	}
	handStrengths := []card.HandStrength{}
	hands := make([][]card.Card, playerNumber)
	hands[0] = slices.Clone(simCommunity)
	hands[0] = append(hands[0], robot.Hand[0])
	hands[0] = append(hands[0], robot.Hand[1])
	combination, err := card.IdentifyCombinations(0, hands[0])
	if err != nil {
		return false, err
	}
	handStrengths = append(handStrengths, combination)
	for i := 1; i < playerNumber; i++ {
		hands[i] = slices.Clone(simCommunity)
		hands[i] = append(hands[i], simulateDeal(deck, &cardsUsed, robot.Hand))
		hands[i] = append(hands[i], simulateDeal(deck, &cardsUsed, robot.Hand))
		combination, err := card.IdentifyCombinations(i, hands[i])
		if err != nil {
			return false, err
		}
		handStrengths = append(handStrengths, combination)
	}
	winner := DetermineWinners(handStrengths)
	return slices.Contains(winner, 0), nil
}

func simulateDeal(deck card.Deck, cardsUsed *int, hand [2]card.Card) card.Card {
	drawnCard := deck[*cardsUsed]
	*cardsUsed++
	if drawnCard == hand[0] || drawnCard == hand[1] {
		return simulateDeal(deck, cardsUsed, hand)
	}
	return drawnCard
}
