package game

import (
	"fmt"
	"slices"
)

func PrintState(deal Deal) {
	fmt.Print("Community cards: ")
	for _, c := range deal.community {
		fmt.Print(c.Read())
		fmt.Print(" ")
	}
	fmt.Println(" ")
	human := deal.players[slices.IndexFunc(deal.players, func(p Player) bool { return p.id == deal.humanPlayerId })]
	humanHand1, humanHand2 := ReadHand(human)
	fmt.Printf("Your hand: %s %s", humanHand1, humanHand2)
	fmt.Println(" ")
	fmt.Printf("Pot: %d", deal.pot)
	fmt.Println(" ")
	fmt.Println(" ")
	for i, p := range deal.players {
		fmt.Print(p.Name)
		if i == deal.dealerId {
			fmt.Print(", dealer")
		}
		if p.HasFolded {
			fmt.Print(", folded")
		}
		fmt.Println(" ")
		fmt.Printf("%d chips, %d bet", p.Chips, p.Bet)
		fmt.Println(" ")
	}
	fmt.Println(" ")
}
