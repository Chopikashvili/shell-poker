package game

import (
	"fmt"
	"slices"

	"github.com/Chopikashvili/shell-poker/general"
	"github.com/Chopikashvili/shell-poker/ux"
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
			fmt.Println("You lost all your chips.")
			fmt.Println("Run the command to play again.")
			break
		} else {
			RunDeal(&deal)
		}
		for i, p := range inst.players {
			if p.Chips != 0 {
				inst.players[i] = deal.players[slices.IndexFunc(deal.players, func(p Player) bool { return p.id == i })]
				inst.players[i].HasFolded = false
				inst.players[i].HasRaised = false
				if slices.ContainsFunc(deal.Winners, func(w Player) bool { return w.Name == p.Name }) {
					inst.players[i].Chips += deal.pot/len(deal.Winners) - inst.players[i].Bet
				} else if p.Chips != 0 {
					inst.players[i].Chips = inst.players[i].Chips - inst.players[i].Bet
				}
			}
		}
		inst.currentDealer++
		if inst.currentDealer == inst.playerNumber {
			inst.currentDealer = 0
		}
	}
}
