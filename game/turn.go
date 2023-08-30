package game

import (
	"chopikashvili/shellpoker/ux"
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
)

func (p *Player) Turn(bets []int) error {
	canBet := !p.HasFolded && !(p.Bet == p.Chips)
	if p.level == 0 || !canBet {
		amount := slices.Max(bets)
		action := ""
		var opt = []string{"call", "raise", "fold"}
		var sel = &survey.Select{Message: "What do you do?", Options: opt}
		err := survey.AskOne(sel, &action, survey.WithIcons(ux.SurveySettings))
		if err != nil {
			return err
		}
		switch action {
		case "call":
			if p.Chips < amount {
				p.callOrRaise(p.Chips) //Could implement the side pot]
			}
			p.callOrRaise(amount)
		case "raise":
			raiseAmount := 0
			survey.AskOne(&survey.Input{Message: "To how much?"}, &raiseAmount, survey.WithIcons(ux.SurveySettings))
			canRaise := raiseAmount > p.Bet && raiseAmount <= p.Chips
			if canRaise {
				p.callOrRaise(raiseAmount)
			} else {
				fmt.Println("Can't raise to that much")
				p.Turn(bets)
			}
		case "fold":
			p.fold()
		}
	} else if canBet {
	}
	return nil
}

func (p *Player) callOrRaise(amount int) {
	p.Bet = amount
}

func (p *Player) fold() {
	p.HasFolded = true
}
