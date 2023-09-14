package game

import (
	"chopikashvili/shellpoker/card"
	"chopikashvili/shellpoker/ux"
	"math/rand"
)

// Initializes game object with number of players "pnum" and level of bots "level".
func BuildGame(set ux.Settings) (GameInstance, error) {
	/*if pnum <= 2 {
		return GameInstance{}, errors.New("Not enough players!")
	} else if pnum >= 10 {
		return GameInstance{}, errors.New("Too many players!")
	} else if level < 1 || level > 10 {
		return GameInstance{}, errors.New("Wrong bot level!")
	}*/
	return buildGame(set)
}

func buildGame(set ux.Settings) (GameInstance, error) {
	game := GameInstance{}
	d, e := card.BuildDeck(set.Letters)
	if e != nil {
		return GameInstance{}, e
	}
	game.gameDeck = d
	game.humanPlayerId = rand.Intn(set.PlayerNumber)
	names, e := GenerateName(set.PlayerNumber)
	if e != nil {
		return GameInstance{}, e
	}
	for i := 0; i < set.PlayerNumber; i++ {
		if i == game.humanPlayerId {
			game.players = append(game.players, Player{Name: "You", Chips: 500, level: 0, id: i}) //initializes a human player object. The hand will be dealt later
		} else {
			game.players = append(game.players, Player{Name: names[i], Chips: 500, level: set.BotLevel, id: i}) //initializes a robot player object with the given level. The hand will be dealt later
		}
	}
	game.currentDeal = 0
	game.currentDealer = rand.Intn(set.PlayerNumber)
	return game, nil
}
