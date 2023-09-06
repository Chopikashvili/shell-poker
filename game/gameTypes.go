package game

import (
	"chopikashvili/shellpoker/card"
)

type GameInstance struct {
	gameDeck      card.Deck //Deck used for game.
	playerNumber  int       //Number of players
	players       []Player  //Info about the players.
	humanPlayerId int       //The ID of the human player, counted "clockwise".
	currentDeal   int       //Number of the current deal.
	currentDealer int       //The ID of the current dealer. The first to bet is 3 places clockwise from the dealer.
	chips         []int     //Tracks how many chips each player has
}

type Deal struct {
	game      GameInstance //The game the deal is part of.
	players   []Player     //The players remaining in the deal
	dealDeck  card.Deck    //Deck used for deal.
	cardsUsed int          //How many cards were dealt to players or to the community.
	community []card.Card  //Community cards.
	bets      []int        //Tracks bets.
	pot       int          //The money in the pot.
	dealerId  int          //Counted clockwise among players still in.
	state     string       //State of the game. May be unnecessary.
	isWon     bool         //Whether the deal has ended.
	Winners   []Player     //Who won the deal
}

type Player struct {
	Name      string       //Name of the player
	id        int          //The ID of the player.
	order     int          //The order of the player from the player with ID 0
	Hand      [2]card.Card //The hole cards of the player.
	Chips     int          //How many chips the player has.
	Bet       int          //The player's bet
	HasFolded bool         //Whether the player has folded in a particular deal.
	level     int          //For robots, the level they play on. For the human player, 0.
}

func ReadHand(pl Player) (string, string) {
	if pl.level == 0 {
		return pl.Hand[0].Read(), pl.Hand[1].Read()
	}
	return "??", "??"
}

func (p Player) GetId() int {
	return p.id
}

func (p Player) GetLevel() int {
	return p.level
}
