package game

import (
	"chopikashvili/shellpoker/card"
	"chopikashvili/shellpoker/general"
	"math/rand"
	"slices"
)

// Initializes deal object
func BuildDeal(g GameInstance) (Deal, error) {
	return buildDeal(g)
}

func buildDeal(g GameInstance) (Deal, error) {
	deal := Deal{}
	deal.game = g
	deal.players = general.Filter(g.players, func(p Player) bool { return p.Chips != 0 })
	for i, p := range deal.players {
		p.HasFolded = false
		p.order = i
		p.deal = &deal
	}
	deal.dealDeck = card.Deck{}
	for i := 0; i < 52; i++ {
		deal.dealDeck[i] = g.gameDeck[i]
	}
	rand.Shuffle(52, func(i, j int) { deal.dealDeck[i], deal.dealDeck[j] = deal.dealDeck[j], deal.dealDeck[i] })
	deal.community = make([]card.Card, 0)
	//determines current dealer, looping through all players to find the first available
	dealerUpNext := g.currentDealer
	for i := dealerUpNext; i > dealerUpNext-g.playerNumber; i-- {
		index := slices.IndexFunc(deal.players, func(p Player) bool { return p.GetId() == i || p.GetId() == i+g.playerNumber })
		if index != -1 {
			deal.dealerId = i
			break
		}
	}
	//determines big blind and small blind
	for i := 0; i < len(deal.players); i++ {
		if i == deal.dealerId+1 || i == deal.dealerId+1-len(deal.players) {
			deal.bets = append(deal.bets, 50)
		} else if i == deal.dealerId+2 || i == deal.dealerId+2-len(deal.players) {
			deal.bets = append(deal.bets, 100)
		} else {
			deal.bets = append(deal.bets, 0)
		}
	}
	deal.cardsUsed = 0
	deal.state = "before betting"
	deal.pot = 150
	for _, p := range deal.game.players {
		p.deal = deal
	}
	return deal, nil
}
