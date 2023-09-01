package game

import (
	"encoding/json"
	"math/rand"
	"os"
	"slices"
)

func GenerateName(amount int) ([]string, error) {
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
