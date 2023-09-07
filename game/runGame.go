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
			fmt.Printf("The game is over. %s won.", winner.Name)
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
		i := 0
		for true {
			p := deal.players[i]
			p.Turn(deal)
			deal.pot = general.Sum(deal.bets)
			deal.Winners = deal.CheckWinner()
			if deal.Winners[0].Name != "" {
				fmt.Printf("%s is the last player standing and wins %d chips", deal.Winners[0].Name, deal.pot)
				break
			}
			i++
			if i == len(deal.players) {
				i = 0
			}
		}
		if deal.Winners[0].Name != "" {
			break
		} else {
			PrintState(*deal)
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
						card1, card2 := ReadHand(p)
						fmt.Printf("%s's hand: %s %s; %s", p.Name, card1, card2, hand.CombName)
						fmt.Println(" ")
					}
				}
				winners := DetermineWinners(hands)
				deal.Winners = general.Filter(deal.players, func(p Player) bool { return slices.Contains(winners, p.id) })
				for _, p := range deal.Winners {
					fmt.Printf("%s wins %d chips!", p.Name, deal.pot/len(deal.Winners))
				}
			}
		}
	}
}
