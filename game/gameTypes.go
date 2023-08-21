package game

import "chopikashvili/shellpoker/card"

type GameInstance struct {
	gameDeck       card.Deck
	playerNumber   int
	players        []Betmaker
	humanPlayerId  int //Counted clockwise.
	currentDeal    int
	startingDealer int
}

type Deal struct {
	game      GameInstance
	dealDeck  card.Deck
	cardsUsed int
	community []card.Card //Community cards
	bets      []int       //Tracks bets.
	dealerId  int         //Counted clockwise.
}

type Human struct {
	id    int
	hand  [2]card.Card
	chips int
}

type Robot struct {
	id    int
	hand  [2]card.Card
	chips int
	level int
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
