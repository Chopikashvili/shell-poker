package game

import (
	"chopikashvili/shellpoker/card"
	"chopikashvili/shellpoker/general"
	"chopikashvili/shellpoker/ux"
	"fmt"
	"slices"
	"time"
)

func RunGame(set ux.Settings) {
	inst, err := BuildGame(set)
	general.Check(err)
	fmt.Print("Players: ")
	for _, p := range inst.players {
		fmt.Printf("%s ", p.Name)
	}
	fmt.Println(" ")
	for true {
		deal, err := BuildDeal(inst)
		general.Check(err)
		if len(deal.players) <= 1 {
			winner := deal.players[0]
			fmt.Printf("The game is over. %s won.", winner.Name)
			break
		} else if slices.IndexFunc(deal.players, func(p Player) bool { return p.id == deal.humanPlayerId }) == -1 {
			fmt.Println("You lost all your money.")
			fmt.Println("Enter 'go run shellpoker' to play again.")
			break
		} else {
			RunDeal(&deal)
		}
		for i, p := range inst.players {
			inst.players[i] = deal.players[i]
			inst.players[i].HasFolded = false
			if slices.ContainsFunc(deal.Winners, func(w Player) bool { return w.Name == p.Name }) {
				inst.players[i].Chips += deal.pot/len(deal.Winners) - inst.players[i].Bet
			} else if p.Chips != 0 {
				inst.players[i].Chips = inst.players[i].Chips - inst.players[i].Bet
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
	PrintState(*deal)
	for i := 0; i < 4; i++ {
		counter := deal.dealerId + 3
		if counter >= len(deal.players) {
			counter = counter - len(deal.players)
		}
		underTheGun := counter
		turns := 0
		for true {
			deal.players[counter].Turn(deal)
			time.Sleep(10 ^ 9)
			for j, p := range deal.players {
				deal.bets[j] = p.Bet
			}
			deal.pot = general.Sum(deal.bets)
			deal.Winners = deal.CheckWinner()
			if len(deal.Winners) != 0 {
				fmt.Printf("%s was the last player standing and won %d chips.", deal.Winners[0].Name, deal.pot)
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
						fmt.Printf("%s's hand: %s %s; %s", p.Name, card1, card2, hand.CombName)
						fmt.Println(" ")
					}
				}
				winners := DetermineWinners(hands)
				deal.Winners = general.Filter(deal.players, func(p Player) bool { return slices.Contains(winners, p.id) })
				for _, p := range deal.Winners {
					fmt.Printf("%s won %d chips!", p.Name, deal.pot/len(deal.Winners))
				}
			}
		}
	}
}

func CheckEndOfTurn(players []Player, underTheGun int, turns *int) bool {
	max := 0
	flag := true
	i, j := underTheGun, underTheGun-1
	if j == -1 {
		j = len(players) - 1
	}
	for true {
		p := players[i]
		if !p.HasFolded && p.Bet < max {
			flag = false
			break
		} else if !p.HasFolded && p.Bet > max {
			max = p.Bet
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
	}
	return flag
}
