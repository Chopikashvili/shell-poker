package game

import (
	"chopikashvili/shellpoker/card"
	"chopikashvili/shellpoker/general"
	"chopikashvili/shellpoker/ux"
	"fmt"
	"slices"
)

func RunGame(set ux.Settings) {
	game, err := BuildGame(set)
	general.Check(err)
	for true {
		deal, err := BuildDeal(game)
		general.Check(err)
		if len(deal.players) <= 1 {
			winner := deal.players[0]
			fmt.Printf("The game is over. %v won.", winner.Name)
			break
		} else {
			RunDeal(&deal)
		}
		for _, p := range game.players {
			if slices.ContainsFunc(deal.Winners, func(w Player) bool { return w.Name == p.Name }) {
				p.Chips += deal.pot / len(deal.Winners)
			} else if p.Chips != 0 {
				p.Chips -= p.Bet
			}
		}
	}
}

func (d Deal) CheckWinner() []Player {
	players := general.Filter(d.players, func(p Player) bool { return !p.HasFolded })
	if len(players) == 1 {
		return []Player{players[0]}
	}
	return []Player{}
}

func RunDeal(deal *Deal) {
	for i := 0; i < 4; i++ {
		for _, p := range deal.players {
			p.Turn(deal)
			deal.pot = general.Sum(deal.bets)
			deal.Winners = deal.CheckWinner()
			if deal.Winners[0].Name != "" {
				fmt.Printf("%v is the last player standing and wins %v chips", deal.Winners[0].Name, deal.pot)
				break
			}
		}
		if deal.Winners[0].Name != "" {
			break
		} else {
			switch i {
			case 0:
				for j := 0; j < 3; j++ {
					deal.community = append(deal.community, deal.dealDeck.Deal(&deal.cardsUsed))
				}
			case 1:
			case 2:
				deal.community = append(deal.community, deal.dealDeck.Deal(&deal.cardsUsed))
			case 3:
				hands := []card.HandStrength{}
				for _, p := range deal.players {
					if !p.HasFolded {
						cardSlice := slices.Clone(deal.community)
						cardSlice = append(cardSlice, p.Hand[0])
						cardSlice = append(cardSlice, p.Hand[1])
						hand, err := card.IdentifyCombinations(p.GetId(), cardSlice)
						general.Check(err)
						hands = append(hands, hand)
					}
				}
				winners := DetermineWinners(hands)
				deal.Winners = general.Filter(deal.players, func(p Player) bool { return slices.Contains(winners, p.id) })
			}
		}
	}
}
