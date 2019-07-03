package assets

import (
	"fmt"
)

// struct to save the rated value and the amount the item has been rated
type dev struct {
	difference  float64
	ratedAmount int
}

// function to create hash map with values
func CreateUserItemRatingsTable() (userRatings map[string]map[string]float64, items []string) {
	userRatings = make(map[string]map[string]float64)
	userRatings["1"] = map[string]float64{}
	userRatings["1"]["103"] = 4.0
	userRatings["1"]["106"] = 3.0
	userRatings["1"]["109"] = 4.0
	userRatings["2"] = map[string]float64{}
	userRatings["2"]["103"] = 5.0
	userRatings["2"]["106"] = 2.0
	userRatings["3"] = map[string]float64{}
	userRatings["3"]["106"] = 3.5
	userRatings["3"]["109"] = 4.0
	userRatings["4"] = map[string]float64{}
	userRatings["4"]["103"] = 5.0
	userRatings["4"]["109"] = 3.0

	items = append(items, "103", "106", "109")

	return userRatings, items
}

// the one slope function
func OneSlope(userRatings map[string]map[string]float64, items []string) {
	// create deviations for all the items users have rated
	deviations := computeDeviations(userRatings, items)
	//printDeviationsOneSlope(deviations, items)

	// loop over every user and check if there aren't empty ratings to predict them.
	for _, value := range userRatings {
		for _, v := range items {
			if _, ok := value[v]; ok {
				continue
			} else {
				// when the value doesn't exist predict and add value to userRatings
				predRating := prediction(value, deviations, v)
				value[v] = predRating
			}
		}
	}
	// predicted ratings results for all the users
	fmt.Println("end result one slope", userRatings)
}

// one slope compute deviation function
func computeDeviations(DevIJ map[string]map[string]float64, items []string) map[string]map[string]dev {
	devTable := make(map[string]map[string]dev)

	// loop over all rated users
	for _, userRatings := range DevIJ {
		i := 0
		// get items of user
		for i < len(items) {
			// check if the item exist
			if _, ok := userRatings[items[i]]; ok {
				j := 0
				// create item when not exist in devTable
				if devTable[items[i]] == nil {
					devTable[items[i]] = map[string]dev{}
				}

				// check if the combination of items exists
				for j < len(items) {
					if i != j {
						if _, ok := userRatings[items[j]]; ok {
							// create item when not exist in item devTable
							if devTable[items[i]][items[j]] == (dev{}) {
								devTable[items[i]][items[j]] = dev{userRatings[items[i]] - userRatings[items[j]], 1}

								// update values when already exists.
							} else {
								newDev := dev{(userRatings[items[i]] - userRatings[items[j]]) + devTable[items[i]][items[j]].difference, devTable[items[i]][items[j]].ratedAmount + 1}
								devTable[items[i]][items[j]] = newDev
							}
						}
					}
					// add j to go to the next item
					j++
				}
			}
			// add i to go to the next item
			i++
		}
	}

	// devide the difference bij total amount rated
	for k, v := range devTable {
		for key, value := range v {
			newVal := dev{value.difference / float64(value.ratedAmount), value.ratedAmount}
			devTable[k][key] = newVal
		}
	}

	return devTable
}

// one slope predictions function
func prediction(userRatings map[string]float64, DevIJ map[string]map[string]dev, itemID string) float64 {
	// result of the combination and how many times are the combinations of items rated (denominator)
	result := 0.0
	denominator := 0

	//loop over list
	for k, v := range userRatings {
		ratedItem := DevIJ[k][itemID]
		result += (v + ratedItem.difference) * float64(ratedItem.ratedAmount)
		denominator += ratedItem.ratedAmount
	}

	res := result / float64(denominator)

	return res
}
