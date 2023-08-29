package game

import (
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
)

var sel = &survey.Select{Message: "What do you do?", Options: opt}
var opt = []string{"call", "raise", "all in", "fold"}

func (p Player) Turn(bets []int) {
	if p.level == 0 || !p.HasFolded {
		action := ""
		survey.AskOne(sel, &action)
		amount := slices.Max(bets)
		switch action {
		case "call":
			p.call(amount)
		case "raise":
			raiseAmount := 0
			survey.AskOne(&survey.Input{Message: "By how much?"}, &raiseAmount)
			canRaise := raiseAmount > 0 && raiseAmount <= p.Chips-p.Bet
			if canRaise {
				p.raise(amount, raiseAmount)
			} else {
				fmt.Println("Can't raise by that much")
				p.Turn(bets)
			}
		case "all in":
			p.raise(amount, p.Chips-amount)
		case "fold":
			p.fold()
		}
	}
}

func (p Player) call(amount int) {
	p.Bet = amount
}

func (p Player) raise(amount int, r int) {
	p.Bet = amount + r
}

func (p Player) fold() {
	p.HasFolded = true
}
