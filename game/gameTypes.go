package game

import "chopikashvili/shellpoker/card"

type GameInstance struct {
	gameDeck       card.Deck  //Deck used for game.
	playerNumber   int        //Number of players
	players        []Betmaker //Info about the players.
	humanPlayerId  int        //The ID of the human player, counted "clockwise".
	currentDeal    int        //Number of the current deal
	startingDealer int        //The ID of the starting dealer. The first to bet is 3 places clockwise from the dealer.
}

type Deal struct {
	game      GameInstance //The game the deal is part of.
	dealDeck  card.Deck    //Deck used for deal (might remove as part of optimization)
	cardsUsed int          //How many cards were dealt to players or to the community.
	community []card.Card  //Community cards.
	bets      []int        //Tracks bets.
	dealerId  int          //Counted clockwise.
}

type Human struct {
	id    int          //The ID of the player.
	hand  [2]card.Card //The hole cards of the player.
	chips int          //How many chips the player has.
}

type Robot struct {
	id    int          //The ID of the player.
	hand  [2]card.Card //The hole cards of the player.
	chips int          //How many chips the player has.
	level int          //For robots, the level they play on.
}

type Betmaker interface {
	GetChips() int
	GetId() int
}

func (h Human) GetChips() int {
	return h.chips
}

func (r Robot) GetChips() int {
	return r.chips
}

func (h Human) GetId() int {
	return h.id
}

func (r Robot) GetId() int {
	return r.id
}
