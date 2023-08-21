package game

import "chopikashvili/shellpoker/card"

type GameInstance struct {
	gameDeck       card.Deck
	playerNumber   int
	players        []any
	humanPlayerId  int //Counted clockwise.
	currentDeal    int
	startingDealer int
	minimumBet     int
}

type Deal struct {
	game      GameInstance
	dealDeck  card.Deck
	cardsUsed int
	community []card.Card
	bets      []int
	dealerId  int //Counted clockwise.
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
