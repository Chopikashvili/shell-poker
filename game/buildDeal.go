package game

import (
	"chopikashvili/shellpoker/card"
	"chopikashvili/shellpoker/general"
)

// Initializes deal object
func BuildDeal(g GameInstance) (Deal, error) {
	return buildDeal(g)
}

func buildDeal(g GameInstance) (Deal, error) {
	deal := Deal{}
	deal.dealDeck = card.Deck{}
	for i := 0; i < 52; i++ {
		deal.dealDeck[i] = g.gameDeck[i]
	}
	deal.dealDeck.Shuffle()
	deal.cardsUsed = 0
	deal.players = general.Filter(g.players, func(p Player) bool { return p.Chips != 0 })
	for i, p := range deal.players {
		p.HasFolded = false
		p.HasRaised = false
		p.order = i
	}
	deal.DealHands()
	deal.community = make([]card.Card, 0)
	//determines current dealer
	deal.dealerId = getDealerId(deal, g)
	if len(deal.players) == 2 {
		deal.setBetsTwoPlayers()
	} else {
		deal.setBets()
	}
	deal.state = "before betting"
	deal.pot = 75
	deal.humanPlayerId = g.humanPlayerId
	return deal, nil
}
