package game

func (d Deal) DealHands() error {
	for _, player := range d.game.players {
		if player.Chips != 0 {
			var e1 error
			var e2 error
			player.Hand[0], e1 = d.game.gameDeck.Deal()
			if e1 != nil {
				return e1
			}
			player.Hand[1], e2 = d.game.gameDeck.Deal()
			if e2 != nil {
				return e2
			}
		}
	}
	return nil
}
