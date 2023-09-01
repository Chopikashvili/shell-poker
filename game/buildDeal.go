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
	}
	deal.dealDeck = card.Deck{}
	for i := 0; i < 52; i++ {
		deal.dealDeck[i] = g.gameDeck[i]
	}
	rand.Shuffle(52, func(i, j int) { deal.dealDeck[i], deal.dealDeck[j] = deal.dealDeck[j], deal.dealDeck[i] })
	deal.community = make([]card.Card, 0)
	//determines current dealer
	dealerId := slices.IndexFunc(deal.players, func(p Player) bool { return p.id == g.currentDealer }) - 1
	if dealerId == -1 {
		deal.dealerId = slices.MaxFunc(deal.players, func(a, b Player) int { return a.id - b.id }).id
	} else {
		deal.dealerId = dealerId
	}
	//determines big blind and small blind
	for i, p := range deal.players {
		if i == deal.dealerId+1 || i == deal.dealerId+1-len(deal.players) {
			if p.Chips < 25 {
				deal.bets = append(deal.bets, p.Chips)
				p.Bet = p.Chips
			} else {
				deal.bets = append(deal.bets, 25)
				p.Bet = 25
			}
		} else if i == deal.dealerId+2 || i == deal.dealerId+2-len(deal.players) {
			if p.Chips < 50 {
				deal.bets = append(deal.bets, p.Chips)
				p.Bet = p.Chips
			} else {
				deal.bets = append(deal.bets, 50)
				p.Bet = 50
			}
		} else {
			deal.bets = append(deal.bets, 0)
		}
	}
	deal.cardsUsed = 0
	deal.state = "before betting"
	deal.pot = 75
	return deal, nil
}
