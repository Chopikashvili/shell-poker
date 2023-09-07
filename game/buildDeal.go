package game

import (
	"chopikashvili/shellpoker/card"
	"chopikashvili/shellpoker/general"
	"slices"
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
	deal.dealDeck.Shuffle()
	deal.cardsUsed = 0
	deal.players = general.Filter(g.players, func(p Player) bool { return p.Chips != 0 })
	for i, p := range deal.players {
		p.HasFolded = false
		p.order = i
	}
	deal.DealHands()
	deal.community = make([]card.Card, 0)
	//determines current dealer
	dealerId := slices.IndexFunc(deal.players, func(p Player) bool { return p.id == g.currentDealer }) - 1
	if dealerId == -1 {
		deal.dealerId = len(deal.players) - 1
	} else {
		deal.dealerId = dealerId
	}
	deal.setBets()
	deal.state = "before betting"
	deal.pot = 75
	return deal, nil
}

func (d *Deal) setBets() {
	for i, p := range d.players {
		if i == d.dealerId+1 || i == d.dealerId+1-len(d.players) {
			if p.Chips < 25 {
				d.bets = append(d.bets, p.Chips)
				p.Bet = p.Chips
			} else {
				d.bets = append(d.bets, 25)
				p.Bet = 25
			}
		} else if i == d.dealerId+2 || i == d.dealerId+2-len(d.players) {
			if p.Chips < 50 {
				d.bets = append(d.bets, p.Chips)
				p.Bet = p.Chips
			} else {
				d.bets = append(d.bets, 50)
				p.Bet = 50
			}
		} else {
			d.bets = append(d.bets, 0)
		}
	}
}
