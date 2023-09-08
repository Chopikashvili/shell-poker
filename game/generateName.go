package game

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"slices"
)

// Generates robot names. Should only be called from the game package.
func GenerateName(amount int) ([]string, error) {
	if amount < 1 || amount > 100 {
		return []string{}, errors.New("Incorrect name amount")
	}
	names := []string{}
	file, err := os.ReadFile("game/names.json")
	if err != nil {
		return []string{}, err
	}
	err = json.Unmarshal(file, &names)
	if err != nil {
		return []string{}, err
	}
	selectedNames := make([]string, amount)
	for i := 0; i < amount; i++ {
		rn := rand.Intn(len(names))
		selectedNames[i] = names[rn]
		slices.Delete(names, rn, rn+1)
	}
	return selectedNames, nil
}
