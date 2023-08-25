package card

//Function for comparing hands. Returns a positive value if hand a is stronger, a negative value if hand b is stronger, and 0 if both are equal
func CompareHands(a, b HandStrength) int {
	if a.CombStrength == b.CombStrength {
		for i := 0; i < 5; i++ {
			if a.OrderedCardValues[i]-b.OrderedCardValues[i] != 0 {
				return a.OrderedCardValues[i] - b.OrderedCardValues[i]
			}
		}
		return 0
	}
	return a.CombStrength - b.CombStrength
}
