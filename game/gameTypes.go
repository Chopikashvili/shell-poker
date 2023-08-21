package game

import "chopikashvili/shellpoker/card"

type GameInstance struct {
	gameDeck      card.Deck
	playerNumber  int
	players       []any
	humanPlayerId int
	currentDeal   int
}

type Deal struct {
	game      GameInstance
	cardsUsed int
	community [5]card.Card
	bets      []int
	dealerId  int
}

type Human struct {
	hand  [2]card.Card
	chips int
}

type Robot struct {
	hand  [2]card.Card
	chips int
	level int
}
