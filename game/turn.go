package game

import (
	"chopikashvili/shellpoker/ux"
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
)

func (p *Player) Turn(deal *Deal) error {
	canBet := !p.HasFolded && !(p.Bet == p.Chips)
	amount := slices.Max(deal.bets)
	if p.level == 0 && canBet {
		action := ""
		var opt = []string{"call"}
		if p.Chips <= amount {
			opt = append(opt, "fold")
		} else {
			opt = append(opt, "raise", "fold")
		}
		var sel = &survey.Select{Message: "What do you do?", Options: opt}
		err := survey.AskOne(sel, &action, survey.WithIcons(ux.SurveySettings))
		if err != nil {
			return err
		}
		switch action {
		case "call":
			p.call(amount)
		case "raise":
			raiseAmount := 0
			survey.AskOne(&survey.Input{Message: "To how much?"}, &raiseAmount, survey.WithIcons(ux.SurveySettings))
			p.humanRaise(amount, raiseAmount, deal)

		case "fold":
			p.HasFolded = true
			fmt.Printf("%s folded.", p.Name)
		}
	} else if canBet {
		action, err := RobotTurn(*p, len(deal.players), deal.community)
		if err != nil {
			return err
		}
		switch action {
		case "call":
			p.call(amount)
			fmt.Printf("%s chose to call.", p.Name)
		case "raise":
			raiseAmount := amount + 50
			p.robotRaise(raiseAmount)
		case "fold":
			p.HasFolded = true
			fmt.Printf("%s folded.", p.Name)
		}
	}
	return nil
}

func (p *Player) call(amount int) {
	if p.Chips < amount {
		p.Bet = p.Chips
	} else {
		p.Bet = amount
	}
	fmt.Printf("%s chose to call.", p.Name)
}

func (p *Player) humanRaise(amount, raiseAmount int, deal *Deal) {
	canRaise := raiseAmount > amount && raiseAmount <= p.Chips
	if canRaise {
		p.Bet = raiseAmount
		fmt.Printf("%s raised to %d.", p.Name, raiseAmount)
	} else {
		fmt.Println("Can't raise to that much")
		p.Turn(deal)
	}
}

func (p *Player) robotRaise(raiseAmount int) {
	if raiseAmount < p.Chips {
		p.Bet = raiseAmount
		fmt.Printf("%s raised to %d.", p.Name, raiseAmount)
	} else {
		p.Bet = p.Chips
		fmt.Printf("%s raised to %d.", p.Name, p.Chips)
	}
}
