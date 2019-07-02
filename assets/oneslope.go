package assets

import (
	"fmt"
)

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

func OneSlope(userRatings map[string]map[string]float64, items []string) {
	// create deviations for all the items users have rated
	deviations := computeDeviations(userRatings, items)

	// loop over every user and check if there aren't empty ratings to predict them.
	for _, value := range userRatings {
		for _, v := range items {
			if _, ok := value[v]; ok {

			} else {
				// when the value doesn't exist predict and add value to userRatings
				predRating := prediction(value, deviations, v)
				value[v] = predRating
			}
		}
	}
	//todo denominator
	fmt.Println("demo")
}

type dev struct {
	difference  float64
	ratedAmount int
}

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

//predictions
func prediction(userRatings map[string]float64, DevIJ map[string]map[string]dev, itemID string) float64 {

	fmt.Println("dit", userRatings, DevIJ, itemID)
	result := 0.0
	amountRated := 0
	for k, v := range userRatings {
		ratedItem := DevIJ[k][itemID]
		fmt.Println("val", (v+ratedItem.difference)*float64(ratedItem.ratedAmount))
		result += (v + ratedItem.difference) * float64(ratedItem.ratedAmount)
		amountRated += ratedItem.ratedAmount
	}

	res := result / float64(amountRated)

	fmt.Println("dem", res)

	return res
}

//if _, ok := DevIJ[items[i]]; !ok {
//fmt.Println("false", v, k)
//}
