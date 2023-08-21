package game

import (
	"chopikashvili/shellpoker/card"
	"errors"
	"math/rand"
)

func BuildGame(pnum int, level int) (GameInstance, error) {
	if pnum <= 2 {
		return GameInstance{}, errors.New("Not enough players!")
	} else if pnum >= 10 {
		return GameInstance{}, errors.New("Too many players!")
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
	game.gameDeck.Shuffle()
	game.playerNumber = pnum
	game.humanPlayerId = rand.Intn(pnum)
	for i := 0; i < pnum; i++ {
		if i == game.humanPlayerId {
			game.players = append(game.players, Human{chips: 1000}) //initializes a human player object. The hand will be dealt later
		} else {
			game.players = append(game.players, Robot{chips: 1000, level: level}) //initializes a robot player object with the given level. The hand will be dealt later
		}
	}
	game.currentDeal = 1
	return game, nil
}
