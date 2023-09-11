package game

import (
	"chopikashvili/shellpoker/card"
	"chopikashvili/shellpoker/general"
	"math/rand"
)

func RobotTurn(robot Player, playerNumber int, community []card.Card, amount int) (string, error) {
	result := ""
	deck, err := card.BuildDeck(false)
	if err != nil {
		return "", err
	}
	scenarioWins := []bool{}
	for i := 0; i < robot.GetLevel(); i++ {
		scenario, err := simulateGame(robot, deck, playerNumber, community)
		if err != nil {
			return "", err
		}
		scenarioWins = append(scenarioWins, scenario)
	}
	winCount := general.Count(scenarioWins, true)
	if robot.GetLevel() == 1 {
		if winCount == 1 {
			if rand.Intn(1) == 0 {
				result = "call"
			} else {
				result = "raise"
			}
		} else {
			result = "fold"
		}
	} else {
		winPct := float64(winCount) / float64(robot.GetLevel())
		if winPct < (2.0/3.0)/float64(playerNumber) && amount > robot.Bet && !robot.HasRaised {
			result = "fold"
		} else if winPct < (4.0/3.0)/float64(playerNumber) {
			result = "call"
		} else {
			result = "raise"
		}
	}
	return BluffCheck(result), nil
}

func BluffCheck(input string) string {
	bluff := rand.Intn(5)
	if bluff == 0 {
		switch input {
		case "fold":
			return "call"
		case "call":
			return "raise"
		case "raise":
			return "call"
		}
	} else {
		return input
	}
	return ""
}
