package game

import (
	"chopikashvili/shellpoker/card"
	"errors"
	"math/rand"
)

// Initializes game object with number of players "pnum" and level of bots "level".
func BuildGame(pnum int, level int) (GameInstance, error) {
	if pnum <= 2 {
		return GameInstance{}, errors.New("Not enough players!")
	} else if pnum >= 10 {
		return GameInstance{}, errors.New("Too many players!")
	} else if level < 1 || level > 10 {
		return GameInstance{}, errors.New("Wrong bot level!")
	}
	return buildGame(pnum, level)
}

func buildGame(pnum int, level int) (GameInstance, error) {
	game := GameInstance{}
	d, e := card.BuildDeck()
	if e != nil {
		return GameInstance{}, e
	}
	game.gameDeck = d
	game.humanPlayerId = rand.Intn(pnum)
	for i := 0; i < pnum; i++ {
		if i == game.humanPlayerId {
			game.players = append(game.players, Player{Chips: 1000, level: 0}) //initializes a human player object. The hand will be dealt later
		} else {
			game.players = append(game.players, Player{Chips: 1000, level: level}) //initializes a robot player object with the given level. The hand will be dealt later
		}
	}
	game.currentDeal = 0
	game.currentDealer = rand.Intn(pnum)
	return game, nil
}
