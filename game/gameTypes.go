package game

import "chopikashvili/shellpoker/card"

type GameInstance struct {
	gameDeck       card.Deck //Deck used for game.
	playerNumber   int       //Number of players
	players        []Player  //Info about the players.
	humanPlayerId  int       //The ID of the human player, counted "clockwise".
	currentDeal    int       //Number of the current deal. One-indexed.
	startingDealer int       //The ID of the starting dealer. The first to bet is 3 places clockwise from the dealer.
}

type Deal struct {
	game      GameInstance //The game the deal is part of.
	dealDeck  card.Deck    //Deck used for deal (might remove as part of optimization)
	cardsUsed int          //How many cards were dealt to players or to the community.
	community []card.Card  //Community cards.
	bets      []int        //Tracks bets.
	pot       int          //The money in the pot.
	dealerId  int          //Counted clockwise.
	state     string       //State of the game.
}

type Player struct {
	id    int          //The ID of the player.
	Hand  [2]card.Card //The hole cards of the player.
	Chips int          //How many chips the player has.
	level int          //For robots, the level they play on. For the human player, 0.
}

func (p Player) ReadHand() (string, string) {
	if p.level == 0 {
		return p.Hand[0].Read(), p.Hand[1].Read()
	}
	return "??", "??"
}

func (p Player) GetId() int {
	return p.id
}

func (p Player) GetLevel() int {
	return p.level
}
