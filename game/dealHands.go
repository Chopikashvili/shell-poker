package game

func (d *Deal) DealHands() error {
	for i, player := range d.players {
		player.Hand[0] = d.dealDeck.Deal(d.cardsUsed)
		d.cardsUsed++
		player.Hand[1] = d.dealDeck.Deal(d.cardsUsed)
		d.cardsUsed++
		d.players[i] = player
	}
	return nil
}
