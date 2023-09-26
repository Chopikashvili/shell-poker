package game

import (
	"github.com/Chopikashvili/shell-poker/general"
	"github.com/Chopikashvili/shell-poker/ux"
)

func PrepareRunningGame() {
	set, err := ux.ConfigureSettings()
	general.Check(err)
	RunGame(set)
}
