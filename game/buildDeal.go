package game

import (
	"chopikashvili/shellpoker/card"
	"errors"
)

// Initializes deal object
func BuildDeal(g GameInstance) (Deal, error) {
	return buildDeal(g)
}

func buildDeal(g GameInstance) (Deal, error) {
	deal := Deal{}
	deal.game = g
	deal.dealDeck = card.Deck{}
	for i := 0; i < 52; i++ {
		deal.dealDeck[i] = g.gameDeck[i]
	}
	deal.community = make([]card.Card, 0)
	//determines current dealer
	i := g.startingDealer - g.currentDeal
	if i >= 0 {
		deal.dealerId = i
	} else if i+g.playerNumber >= 0 {
		deal.dealerId = i + g.playerNumber
	} else {
		return Deal{}, errors.New("Something went wrong while assigning dealer")
	}
	//determines big blind and small blind
	for i := 0; i < g.playerNumber; i++ {
		if i == deal.dealerId+1 || i == deal.dealerId+1-g.playerNumber {
			deal.bets = append(deal.bets, 50)
		} else if i == deal.dealerId+2 || i == deal.dealerId+2-g.playerNumber {
			deal.bets = append(deal.bets, 100)
		} else {
			deal.bets = append(deal.bets, 0)
		}
	}
	return deal, nil
}

func copySlice[T any](slice []T) []T {
	copy := []T{}
	for _, elem := range slice {
		copy = append(copy, elem)
	}
	return copy
}
