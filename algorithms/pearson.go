package algorithms

import (
	"math"
)

func Pearson(user map[string]float64, otherUser map[string]float64) (distance float64) {
	// Calculates the distance between the same items and adds similarity to see the difference in total ratings
	var A, B1, B2, C1, C2, D1, D2 float64
	lengthItems := 0.0

	// go true every item of the user
	for k := range user {
		// check if the item exist for the other user
		// Only calculate distance when the other user rated the same items
		if otherUserRating, ok := otherUser[k]; ok {
			// Distance multiplied by the power of two
			A += distanceMultipliedBetweenXY(user[k], otherUserRating)
			B1 += user[k]
			B2 += otherUserRating

			C1 += math.Pow(user[k], 2)
			D1 += math.Pow(otherUserRating, 2)
			lengthItems++
		}
	}

	upper := A - ((B1 * B2) / lengthItems)
	C2 = math.Pow(B1, 2) / lengthItems
	D2 = math.Pow(B2, 2) / lengthItems
	under := math.Sqrt(C1-C2) * math.Sqrt(D1-D2)
	return upper / under
}
