package main

import (
	"fmt"
	"hro.projects/INFDAT01-2NEW/assets"
	"sort"
)

type strategyPattern interface {
	Similarity(map[string]map[string]float64, []string, map[string]float64) (map[string]map[string]map[float64]int, error)
	Average(map[string]float64) float64
}

type itemItem struct {
	dataset             *map[string]map[string]float64
	averages            *map[string]float64
	calculateAverage    strategyPattern
	calculateSimilarity strategyPattern
	algorithmName       string
	itemsIDs            *[]string
	similarities        *map[string]map[string]map[float64]int
}

type CosineItem struct {
}

func (cosine CosineItem) Average(user map[string]float64) (distance float64) {
	// Executes the distance between the same items and adds similarity to see the difference in total ratings
	integer := 0.0
	sum := 0.0
	for _, value := range user {
		sum += value
		integer++
	}

	return sum / integer
}

func (uI *itemItem) Average() error {
	tempDataset, err := assets.ReadDataset("files/item-item.txt")
	//tempDataset, err := assets.ReadMovieDataSet("files/movieLens100KUserItems.data")
	if err != nil {
		return fmt.Errorf("database is nog leeg er kan geen recomendatie gedaan worden")
	}
	*uI.dataset = tempDataset

	// Declare new variables to prevent repeatability and add readability.
	var UserItemMap = tempDataset
	usersAverages := make(map[string]float64)
	var itemIDs []string

	// walk true every item from the user
	for _, val := range UserItemMap {
		for k2 := range val {
			if assets.ItemContains(itemIDs, k2) {
				itemIDs = append(itemIDs, k2)
			}
		}
	}

	// walk true every user
	for k, val := range UserItemMap {
		averageRatingUser := uI.calculateAverage.Average(val)
		//usersAverages = append(usersAverages, averageRatingUser)
		usersAverages[k] = averageRatingUser
	}

	// sorting the keys
	usersAverages = assets.SortList(usersAverages)

	sort.Strings(itemIDs)
	*uI.itemsIDs = itemIDs
	*uI.averages = usersAverages

	return nil
}

func (cosine CosineItem) Similarity(dataset map[string]map[string]float64, itemIds []string, userAverages map[string]float64) (cosineAdjustedSimilarity map[string]map[string]map[float64]int, err error) {
	cosineAdjustedSimilarity = make(map[string]map[string]map[float64]int)
	var loopedOver []string

	for _, v := range itemIds {
		if cosineAdjustedSimilarity[v] == nil {
			cosineAdjustedSimilarity[v] = map[string]map[float64]int{}
		}
		for _, value := range itemIds {
			if v != value && assets.ContainsString(loopedOver, value) != true {
				if cosineAdjustedSimilarity[v][value] == nil {
					cosineAdjustedSimilarity[v][value] = map[float64]int{}
				}
			}
		}

		loopedOver = append(loopedOver, v)
	}

	//call the cosineAdjustedSimilarity calculate Function
	cas, err := assets.ACS(cosineAdjustedSimilarity, dataset, itemIds, userAverages)
	if err != nil {
		return cosineAdjustedSimilarity, err
	}

	// return full cosineAdjustedSimilarity
	return cas, nil
}

// Execute interface executed for every method in StrategyPattern
func (uI *itemItem) Similarity() error {
	strategyCalculatedSimilarity, err := uI.calculateSimilarity.Similarity(*uI.dataset, *uI.itemsIDs, *uI.averages)
	if err != nil {
		return fmt.Errorf("error when execute item-item strategypattern")
	}

	uI.similarities = &strategyCalculatedSimilarity

	return nil
}

func main() {
	//variables
	var itemIDS []string
	var itemDataset map[string]map[string]float64
	var itemAverages map[string]float64
	var cosineAlg CosineItem
	newUserID := "6"
	findUnratedItem := "103"
	// create ItemItem with different algorithms at runtime
	cosine := itemItem{calculateSimilarity: cosineAlg, algorithmName: "Cosine distance\n similarity", itemsIDs: &itemIDS, dataset: &itemDataset, calculateAverage: cosineAlg, averages: &itemAverages}
	// get averages of users from dataset
	err := cosine.Average()

	if err != nil {
		fmt.Println("cosine Average", err.Error())
	} else {
		// print result off the user averages
		assets.PrintItemItemAverages(*cosine.averages, "The average rating for all the users:")

		// Step 1 ACS
		// calculate similarity between items
		err := cosine.Similarity()
		if err != nil {
			fmt.Println("Cosine Similarity", err)
		} else {
			assets.PrintItemAlgorithmSimilarities(cosine.similarities, *cosine.itemsIDs, "Cosine adjustment formula:  similarity between all items")
		}

		// create new user with item ratings
		userRatings, err := assets.CreateNewUser("6")

		if err != nil {
			fmt.Println("User ratings", err)
		}

		// Get the lowest and highest rated item from the user and Normalise user ratings.
		var normalizedUserRatings float64
		userLowestRating, userHighestRating := assets.GetUserLowestAndHighestValue(userRatings[newUserID])

		for k, v := range *cosine.similarities {
			if k == findUnratedItem {
				normalizedUserRatings, err = assets.NormalizeUserRatings(v, userRatings[newUserID], userLowestRating, userHighestRating)

				if err != nil {
					fmt.Println("Normalized ratings", err)
				}
			}
		}

		// Denormalize value to suggest the predicted rating for the user
		denormalizePredictedUserRating := assets.DenormalizeValue(normalizedUserRatings, userLowestRating, userHighestRating)

		fmt.Println("\nNormalized user item similarity \n", normalizedUserRatings)
		fmt.Println("Denormalized predicted rating for 103 \n", denormalizePredictedUserRating)
		fmt.Println("\n\n\nOneslope\n")
		// the normalisation could be used to predict the rating for item 103

		// Step 2 ONE SLOPE
		// getUserRatings
		oneSlopeUserRatings, oneSlopeItems := assets.CreateUserItemRatingsTable()

		// compute oneSlope
		assets.OneSlope(oneSlopeUserRatings, oneSlopeItems)
	}
}
