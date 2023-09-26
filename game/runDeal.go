package game

import (
	"chopikashvili/shellpoker/card"
	"chopikashvili/shellpoker/general"
	"fmt"
	"slices"
	"time"
)

func RunDeal(deal *Deal) {
	PrintState(*deal)
	for i := 0; i < 4; i++ {
		counter := 0
		if i == 0 {
			if len(deal.players) == 2 {
				counter = deal.dealerId
			} else {
				counter = deal.dealerId + 3
			}
		} else {
			counter = deal.dealerId + 1
		}
		for counter >= len(deal.players) {
			counter = counter - len(deal.players)
		}
		underTheGun := counter
		turns := 0
		for true {
			deal.players[counter].Turn(deal)
			time.Sleep(time.Second)
			for j, p := range deal.players {
				deal.bets[j] = p.Bet
			}
			deal.pot = general.Sum(deal.bets)
			deal.Winners = deal.CheckWinner()
			if len(deal.Winners) != 0 {
				if deal.Winners[0].Name == "You" {
					fmt.Printf("You were the last player standing and won %d chips.", deal.pot-deal.Winners[0].Bet)
				} else {
					fmt.Printf("%s was the last player standing and won %d chips.", deal.Winners[0].Name, deal.pot-deal.Winners[0].Bet)
				}
				fmt.Println(" ")
				break
			}
			counter++
			if counter == len(deal.players) {
				counter = 0
			}
			if CheckEndOfTurn(deal.players, underTheGun, &turns) {
				break
			}
		}
		if len(deal.Winners) != 0 {
			break
		} else {
			switch i {
			case 0:
				for j := 0; j < 3; j++ {
					deal.community = append(deal.community, deal.dealDeck.Deal(&deal.cardsUsed))
				}
				PrintState(*deal)
			case 1:
				deal.community = append(deal.community, deal.dealDeck.Deal(&deal.cardsUsed))
				PrintState(*deal)
			case 2:
				deal.community = append(deal.community, deal.dealDeck.Deal(&deal.cardsUsed))
				PrintState(*deal)
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
						fmt.Printf("%s: %s %s; %s", p.Name, card1, card2, hand.CombName)
						fmt.Println(" ")
					}
				}
				winners := DetermineWinners(hands)
				deal.Winners = general.Filter(deal.players, func(p Player) bool { return slices.Contains(winners, p.id) })
				for _, p := range deal.Winners {
					fmt.Printf("%s won %d chips!", p.Name, (deal.pot/len(deal.Winners))-p.Bet)
					fmt.Println(" ")
				}
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

func CheckEndOfTurn(players []Player, underTheGun int, turns *int) bool {
	max := 0
	flag := true
	i, j := underTheGun, underTheGun
	if j == -1 {
		j = len(players) - 1
	}
	for _, p := range players {
		if !p.HasFolded && p.Bet > max {
			max = p.Bet
		}
	}
	for true {
		p := players[i]
		if !p.HasFolded && p.Chips != p.Bet && p.Bet < max {
			flag = false
			break
		}
		i++
		if i == len(players) {
			i = 0
		}
		if i == j {
			break
		}
	}
	//this is here so at the start of the deal the check does not immediately fire
	*turns = *turns + 1
	if *turns < len(players) {
		return false
	} else if general.CountFunc(players, func(p Player) bool { return !p.HasFolded && p.Chips != p.Bet }) <= 1 {
		return true
	}
	return flag
}
