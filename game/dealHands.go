package game

func (d Deal) DealHands() error {
	for i, player := range d.game.players {
		if player.Chips != 0 {
			player.Hand[0] = d.dealDeck.Deal(d.cardsUsed)
			d.cardsUsed++
			player.Hand[1] = d.dealDeck.Deal(d.cardsUsed)
			d.cardsUsed++
			d.game.players[i] = player
		}
	}
	return nil
}
