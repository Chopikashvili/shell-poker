package card

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
