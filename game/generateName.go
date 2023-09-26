package game

import (
	"errors"
	"math/rand"
	"slices"
)

// Generates robot names. Should only be called from the game package.
func GenerateName(amount int) ([]string, error) {
	if amount < 1 || amount > 100 {
		return []string{}, errors.New("Incorrect name amount")
	}
	names := GameNames
	selectedNames := make([]string, amount)
	for i := 0; i < amount; i++ {
		rn := rand.Intn(len(names))
		selectedNames[i] = names[rn]
		slices.Delete(names, rn, rn+1)
	}
	return selectedNames, nil
}
