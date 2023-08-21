package game

import (
	"chopikashvili/shellpoker/card"
	"errors"
)

func BuildDeal(g GameInstance) (Deal, error) {
	return buildDeal(g)
}

func buildDeal(g GameInstance) (Deal, error) {
	deal := Deal{}
	deal.game = g
	deck := g.gameDeck
	deck.Shuffle()
	deal.dealDeck = deck
	deal.community = make([]card.Card, 0)
	i := g.startingDealer - g.currentDeal
	if i >= 0 {
		deal.dealerId = i
	} else if i+g.playerNumber >= 0 {
		deal.dealerId = i + g.playerNumber
	} else {
		return Deal{}, errors.New("Something went wrong while assigning dealer")
	}
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
