package game

import (
	"fmt"
	"slices"
)

func PrintState(deal Deal) {
	fmt.Print("Community cards:")
	for _, c := range deal.community {
		fmt.Print(c.Read())
	}
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Printf("Pot: %d", deal.pot)
	fmt.Println(" ")
	fmt.Println(" ")
	for _, p := range deal.players {
		fmt.Println(p.Name)
		fmt.Printf("%d chips, %d bet", p.Chips, p.Bet)
		fmt.Println(" ")
	}
	human := deal.players[slices.IndexFunc(deal.players, func(p Player) bool { return p.id == deal.humanPlayerId })]
	humanHand1, humanHand2 := ReadHand(human)
	fmt.Printf("Your hand: %s %s", humanHand1, humanHand2)
	fmt.Println(" ")
}
