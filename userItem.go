package main

import (
	"fmt"
	"github.com/umpc/go-sortedmap"
	"github.com/umpc/go-sortedmap/asc"
	"hro.projects/INFDAT01-2NEW/algorithms"
	"hro.projects/INFDAT01-2NEW/assets"
	"math"
	"strconv"
)

// struct which contains nearest neighbour information
type nearestNeighbour struct {
	userID         string
	data           assets.Data
	similarity     assets.NewUserItemDataSet
	threshold      float64
	nearest        *map[string]map[string]float64
	userRatedItems []string
}

//
// StrategyPattern on Algorithms
//

type StrategyPattern interface {
	Calculate(map[string]float64, map[string]float64) float64
}

type userItem struct {
	userID        string
	dataset       *map[string]map[string]float64
	calculate     StrategyPattern
	algorithmName string
}

// Euclidean
type Euclidean struct {
}

func (euclidean Euclidean) Calculate(user map[string]float64, otherUser map[string]float64) (distance float64) {
	// Calculates the distance between the same items and adds similarity to see the difference in total ratings
	var distanceDifference = 0.0

	for k := range user {
		// Only calculate distance when the other user rated the same items
		if otherUserRating, ok := otherUser[k]; ok {
			// Distance multiplied by the power of two
			distanceDifference += math.Pow(algorithms.DistanceBetweenXY(user[k], otherUserRating), 2)
		}
	}

	// Get the squared root off the totalDistance and return the similarity distance.
	totalDistance := math.Sqrt(distanceDifference)
	return algorithms.Similarity(totalDistance)
}

// Pearson
type Pearson struct {
}

func (pearson Pearson) Calculate(user map[string]float64, otherUser map[string]float64) (distance float64) {
	// Calculates the distance between the same items and adds similarity to see the difference in total ratings
	var A, B1, B2, C1, C2, D1, D2 float64
	lengthItems := 0.0

	// go true every item of the user
	for k := range user {
		// check if the item exist for the other user
		// Only calculate distance when the other user rated the same items
		if otherUserRating, ok := otherUser[k]; ok {
			// Distance multiplied by the power of two
			A += algorithms.DistanceMultipliedBetweenXY(user[k], otherUserRating)
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

// Pearson
type Cosine struct {
}

func (cosine Cosine) Calculate(user map[string]float64, otherUser map[string]float64) (distance float64) {
	// Calculates the distance between the same items and adds similarity to see the difference in total ratings
	var A, B, C float64

	for k := range user {
		// Only calculate distance when the other user rated the same items
		if otherUserRating, ok := otherUser[k]; ok {
			A += algorithms.DistanceMultipliedBetweenXY(user[k], otherUserRating)
			B += math.Pow(user[k], 2)
			C += math.Pow(otherUserRating, 2)
		}
	}

	upper := A
	under := math.Sqrt(B) * math.Sqrt(C)

	return upper / under
}

// Calculate interface executed for every method in StrategyPattern
func (uI *userItem) Calculate() assets.NewUserItemDataSet {
	uI.userID = "7"
	tempDataset := assets.ReadDataset("files/user-item.txt")
	//tempDataset := assets.ReadMovieDataSet("files/movieLens100KUserItems.data")
	uI.dataset = &tempDataset

	distanceBetweenUserAndUsers := make(map[string]float64)

	// Declare new variables to prevent repeatability and add readability.
	var UserItemMap = *uI.dataset
	var userID = uI.userID

	// walk true every item from the user
	for k := range UserItemMap {
		if userID != k {
			distanceBetweenUsers := uI.calculate.Calculate(UserItemMap[userID], UserItemMap[k])
			distanceBetweenUserAndUsers[k] = distanceBetweenUsers
		}
	}
	return assets.NewUserItemDataSet{UserID: uI.userID, Dataset: distanceBetweenUserAndUsers, AlgorithmName: uI.algorithmName}
}

//
// Nearest Neighbours
//

// neighbours method to calculate nearestNeighbours
func (neighbours *nearestNeighbour) calculate() {
	//todo create algorithm to decide the unique items needed (Mean) for better results
	//todo upgrade threshold if list is full
	//variables
	sim := neighbours.similarity.Dataset
	uniqueUserItems := *neighbours.data.UniqueUserItemRatings
	totalNeighbours := 3
	sortedMap := sortedmap.New(totalNeighbours, asc.Float64)
	var nearestNeighbours map[string]float64
	nearestNeighbours = map[string]float64{}

	// walk true every similarity in the dataset
	for key, value := range sim {
		// only adds user when similarity is higher then the threshold
		if value > neighbours.threshold {
			// only adds user if it has more than 1 unique item rated
			if len(uniqueUserItems[key]) > 1 {
				nearestNeighbours[key] = value
				sortedMap.Insert(key, value)
			}
		}
		// break after finding totalNeighbours
		if len(nearestNeighbours) == totalNeighbours {
			break
		}
	}
	neighboursList := map[string]map[string]float64{}
	// select values > lowerBound and values <= upperBound.
	// loop through the values, in reverse order:
	iterCh, err := sortedMap.BoundedIterCh(true, 0.0, 1.0)
	if err != nil {
		fmt.Println("error while sorting list", err)
	}

	// set neighboursList on order from best item rating
	for rec := range iterCh.Records() {
		key := rec.Key.(string)
		temp := *neighbours.data.UniqueUserItemRatings
		neighboursList[key] = temp[key]
	}

	// set nearestNeighbours
	neighbours.nearest = &neighboursList
}

// neighbours method to predict 'item rating' based on nearest neighbours
func (neighbours *nearestNeighbour) predictUniqueItemRatings() {
	//variables
	ItemsPearsonRanked := map[string]float64{}
	listTimesOfAllItems := map[string]float64{}
	dataset := *neighbours.nearest

	//count the similarity for every unique item.
	for userID, algorithmDistance := range neighbours.similarity.Dataset {
		for key, value := range dataset[userID] {
			// add items to the list to create a specific search later
			predictedRatingForItem := value * algorithmDistance
			ItemsPearsonRanked[key] += predictedRatingForItem
			listTimesOfAllItems[key] += algorithmDistance
		}
	}

	sortedMap := sortedmap.New(len(ItemsPearsonRanked), asc.Float64)

	// divide the sum of the item similarity to predict the users items rating
	for key := range ItemsPearsonRanked {
		totalRating := ItemsPearsonRanked[key]
		timesRated := listTimesOfAllItems[key]
		ItemsPearsonRanked[key] = totalRating / timesRated
		sortedMap.Insert(key, totalRating/timesRated)
	}

	iterCh, err := sortedMap.BoundedIterCh(true, 0.0, 10.0)
	if err != nil {
		fmt.Println("error while sorting list", err)
	}

	fmt.Println("\nPredicted Items result:")
	userPredictedItemRatings := map[string]float64{}
	// needed to return max 10 results
	i := 0
	for rec := range iterCh.Records() {
		i += 1
		integer := strconv.Itoa(i)
		key := rec.Key.(string)
		userPredictedItemRatings[key] = ItemsPearsonRanked[key]

		// needed for result printing reasons
		value := fmt.Sprintf("%.2f", ItemsPearsonRanked[key])
		fmt.Println(integer + " : Item: " + key + " Rating: " + value)
		if i == 20 {
			break
		}
	}
}

// UserItem method to find unique and same ratings as user
func (uI *userItem) findUsersWithMoreUniqueRatings() (assets.Data, []string) {
	// variables for readability
	userID := uI.userID
	userData := *uI.dataset
	userRatings := userData[userID]
	var userRatedItems []string

	// new dataset which contains the unique ratings off other users.
	sameRatingsAsUser := map[string]map[string]float64{}
	datasetWithUniqueRatings := map[string]map[string]float64{}

	// loop over all the userRatings
	for otherUserID, otherRatings := range userData {
		// Skip if the id is the same as user
		if otherUserID != userID {
			// check if the other user has minimal one rating similar as user:
			if assets.Contains(userRatings, otherRatings) {
				// when user doesn't exist in the datasetWithUniqueRatings add user
				if datasetWithUniqueRatings[otherUserID] == nil {
					datasetWithUniqueRatings[otherUserID] = map[string]float64{}
					sameRatingsAsUser[otherUserID] = map[string]float64{}
				}

				// loop over all the ratings to find unique ratings to add tot the new dataset
				for key, value := range otherRatings {
					if _, ok := userRatings[key]; ok {
						sameRatingsAsUser[otherUserID][key] = value
					} else {
						//add key and value to new list
						datasetWithUniqueRatings[otherUserID][key] = value
					}
				}
			}
		} else {
			for key := range otherRatings {

				userRatedItems = append(userRatedItems, key)
			}
		}
	}

	userItemDataset := assets.Data{EqualUserItemRatings: &sameRatingsAsUser, UniqueUserItemRatings: &datasetWithUniqueRatings, AllUserItemRatings: uI.dataset}
	return userItemDataset, userRatedItems
}

func main() {
	// variables
	var list []assets.NewUserItemDataSet
	threshold := 0.8

	// create userItems with different algorithms at runtime
	euclidean := userItem{calculate: Euclidean{}, algorithmName: "Euclidean distance\n similarity"}
	pearson := userItem{calculate: Pearson{}, algorithmName: "Pearson \n coefficient"}
	cosine := userItem{calculate: Cosine{}, algorithmName: "Cosine"}

	// calculate algorithms
	euclideanResult := euclidean.Calculate()
	pearsonResult := pearson.Calculate()
	cosineResult := cosine.Calculate()

	// create a list off all the results
	list = append(list, pearsonResult, euclideanResult, cosineResult)

	// print result off the algorithms
	assets.PrintMultipleAlgorithms(list, "The distance from user "+pearson.userID+" compared with the other users:")

	// PART 3
	// find similar and unique ratings
	userSeven := userItem{pearson.userID, pearson.dataset, Pearson{}, "pearson"}
	equalAndUniqueUserItemRatings, userRatedItems := userSeven.findUsersWithMoreUniqueRatings()
	// print result off the dataset with the same and different ratings.
	assets.PrintsSimilarAndDifferentItems(equalAndUniqueUserItemRatings, "See the same and unique ratings for each user compared with user "+pearson.userID+":")

	// Part 4
	// Nearest Neighbour
	nearestNeighbour := nearestNeighbour{pearson.userID, equalAndUniqueUserItemRatings, pearsonResult, threshold, nearestNeighbour{}.nearest, userRatedItems}
	nearestNeighbour.calculate()

	// Part 5
	nearestNeighbour.predictUniqueItemRatings()
	fmt.Println("End result")
}
