package algorithms

import "math"

func Cosine(user map[string]float64, otherUser map[string]float64) (distance float64){
	// Calculates the distance between the same items and adds similarity to see the difference in total ratings
	var A, B, C float64

	for k := range user {
		// Only calculate distance when the other user rated the same items
		if otherUserRating, ok := otherUser[k]; ok {
			A += distanceMultipliedBetweenXY(user[k], otherUserRating)
			B += math.Pow(user[k], 2)
			C += math.Pow(otherUserRating, 2)
		}
	}

	upper := A
	under := math.Sqrt(B) * math.Sqrt(C)

	return upper / under
}
