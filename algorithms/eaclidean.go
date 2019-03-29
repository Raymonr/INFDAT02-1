package algorithms

import "math"

func EuclideanDistance(user map[string]float64, otherUser map[string]float64) (distance float64){
	// Calculates the distance between the same items and adds similarity to see the difference in total ratings
	var distanceDifference = 0.0

	for k := range user {
		// Only calculate distance when the other user rated the same items
		if otherUserRating, ok := otherUser[k]; ok {
			// Distance multiplied by the power of two
			distanceDifference += math.Pow(distanceBetweenXY(user[k], otherUserRating), 2)
		}
	}

	// Get the squared root off the totalDistance and return the similarity distance.
	totalDistance := math.Sqrt(distanceDifference)
	return similarity(totalDistance)
}
