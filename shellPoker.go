package main

import (
	"chopikashvili/shellpoker/game"
	"chopikashvili/shellpoker/general"
	"chopikashvili/shellpoker/ux"
)

func main() {
	set, err := ux.ConfigureSettings()
	general.Check(err)
	game.RunGame(set)
}
