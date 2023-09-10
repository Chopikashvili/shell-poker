package game

import (
	"chopikashvili/shellpoker/general"
	"chopikashvili/shellpoker/ux"
)

func PrepareRunningGame() {
	set, err := ux.ConfigureSettings()
	general.Check(err)
	RunGame(set)
}
