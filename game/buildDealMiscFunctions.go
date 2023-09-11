package game

import "slices"

func (d *Deal) setBets() {
	for i, p := range d.players {
		if i == d.dealerId+1 || i == d.dealerId+1-len(d.players) {
			if p.Chips < 25 {
				d.bets = append(d.bets, p.Chips)
				d.players[i].Bet = p.Chips
			} else {
				d.bets = append(d.bets, 25)
				d.players[i].Bet = 25
			}
		} else if i == d.dealerId+2 || i == d.dealerId+2-len(d.players) {
			if p.Chips < 50 {
				d.bets = append(d.bets, p.Chips)
				d.players[i].Bet = p.Chips
			} else {
				d.bets = append(d.bets, 50)
				d.players[i].Bet = 50
			}
		} else {
			d.bets = append(d.bets, 0)
			d.players[i].Bet = 0
		}
	}
}

func (d *Deal) setBetsTwoPlayers() {
	for i, p := range d.players {
		if i == d.dealerId {
			if p.Chips < 25 {
				d.bets = append(d.bets, p.Chips)
				d.players[i].Bet = p.Chips
			} else {
				d.bets = append(d.bets, 25)
				d.players[i].Bet = 25
			}
		} else {
			if p.Chips < 50 {
				d.bets = append(d.bets, p.Chips)
				d.players[i].Bet = p.Chips
			} else {
				d.bets = append(d.bets, 50)
				d.players[i].Bet = 50
			}
		}
	}
}

func getDealerId(deal Deal, g GameInstance) int {
	dealerId := slices.IndexFunc(deal.players, func(p Player) bool { return p.id == g.currentDealer }) - 1
	if dealerId == -1 {
		return len(deal.players) - 1
	} else if dealerId == -2 {
		//this section is for if the previous dealer went bankrupt
		for i := g.currentDealer; i >= 0; i-- {
			dealerId := slices.IndexFunc(deal.players, func(p Player) bool { return p.id == i }) - 1
			if dealerId != -2 {
				return dealerId
			}
		}
		for i := g.playerNumber; i > g.currentDealer; i-- {
			dealerId := slices.IndexFunc(deal.players, func(p Player) bool { return p.id == i }) - 1
			if dealerId != -2 {
				return dealerId
			}
		}
	} else {
		return dealerId
	}
	return 0
}
