package ux

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
)

type Settings struct {
	Letters      bool
	PlayerNumber int
	BotLevel     int
}

var qs = []*survey.Question{
	{Name: "letters", Prompt: &survey.Confirm{Message: "Do you see playing card suits: ♠ ♣ ♥ ♦ If not, you probably want to play with letters representing card suits. Do you want to play with letters as suits?"}},
	{Name: "playerNumber", Prompt: &survey.Input{Message: "Next, enter a number from 2 to 10 to determine how many players you want to have in the game."}},
	{Name: "botLevel", Prompt: &survey.Input{Message: "Finally, enter a number from 1 to 10 to set the level of bots you'll be playing"}},
}

func ConfigureSettings() (Settings, error) {
	set := Settings{}
	err := survey.Ask(qs, &set)
	if err != nil {
		return Settings{}, err
	}
	if set.PlayerNumber < 2 || set.PlayerNumber > 10 {
		return Settings{}, errors.New("Wrong number of players!")
	} else if set.BotLevel < 1 || set.BotLevel > 10 {
		return Settings{}, errors.New("Wrong bot level!")
	}
	return set, nil
}
