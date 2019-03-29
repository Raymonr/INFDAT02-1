package main

import (
	"fmt"
	"github.com/umpc/go-sortedmap"
	"github.com/umpc/go-sortedmap/asc"
	"hro.projects/INFDAT01-2NEW/algorithms"
	"hro.projects/INFDAT01-2NEW/assets"
	"strconv"
)

// types struct are the ways to create classes in Go
// capital typed types are public lowercase(camelcase) are private
// all is commented to explain whats happens

// struct which contains dataset and user id
type UserItem struct{
	userID string
	dataset *map[string]map[string]float64
}

// struct which contains nearest neighbour information
type nearestNeighbour struct{
	userID    string
	data      assets.Data
	similarity assets.NewUserItemDataSet
	threshold float64
	nearest 	*map[string]map[string]float64
	userRatedItems []string
}

// UserItem method for the euclidean algorithm with a reference to the heap
func (uI *UserItem) euclidean() assets.NewUserItemDataSet {
	// declare new variables to prevent repeatability and add readability.
	distanceBetweenUserAndUsers := make(map[string]float64)
	var UserItemMap = *uI.dataset
	var userID 			= uI.userID

	// walk true every item from the user
	for k := range UserItemMap {
		if userID != k {
			distanceBetweenUsers := algorithms.EuclideanDistance(UserItemMap[userID], UserItemMap[k])
			distanceBetweenUserAndUsers[k] = distanceBetweenUsers
		}
	}
	// return value with NewUserItemDataSet which contains current user, newDataset and explains the algorithm.
	return assets.NewUserItemDataSet{UserID: uI.userID, Dataset: distanceBetweenUserAndUsers, AlgorithmName: "Euclidean distance\n similarity"}
}

// UserItem method for the pearson algorithm with a reference to the heap
func (uI *UserItem) pearson() assets.NewUserItemDataSet {
	distanceBetweenUserAndUsers := make(map[string]float64)

	// Declare new variables to prevent repeatability and add readability.
	var UserItemMap = *uI.dataset
	var userID 			= uI.userID

	// walk true every item from the user
	for k := range UserItemMap {
		if userID != k {
			distanceBetweenUsers := algorithms.Pearson(UserItemMap[userID], UserItemMap[k])
			distanceBetweenUserAndUsers[k] = distanceBetweenUsers
		}
	}
	// return value with NewUserItemDataSet which contains current user, newDataset and explains the algorithm.
	return assets.NewUserItemDataSet{UserID: uI.userID, Dataset: distanceBetweenUserAndUsers, AlgorithmName: "Pearson \n coefficient"}
}

// UserItem method for the cosine algorithm with a reference to the heap
func (uI *UserItem) cosine() assets.NewUserItemDataSet {
	distanceBetweenUserAndUsers := make(map[string]float64)

	// declare new variables to prevent repeatability and add readability.
	var UserItemMap = *uI.dataset
	var userID 			= uI.userID

	// walk true every item from the user
	for k := range UserItemMap {
		if userID != k {
			// go to algorithms folder and
			distanceBetweenUsers := algorithms.Cosine(UserItemMap[userID], UserItemMap[k])
			distanceBetweenUserAndUsers[k] = distanceBetweenUsers
		}
	}
	// return value with NewUserItemDataSet which contains current user, newDataset and explains the algorithm.
	return assets.NewUserItemDataSet{UserID: uI.userID, Dataset: distanceBetweenUserAndUsers, AlgorithmName: "Cosine"}
}

// UserItem method to find unique and same ratings as user
func (uI *UserItem) findUsersWithMoreUniqueRatings() (assets.Data, []string) {
	// variables for readability
	userID := uI.userID
	userData := *uI.dataset
	userRatings := userData[userID]
	var userRatedItems []string

	// new dataset which contains the unique ratings off other users.
	sameRatingsAsUser :=  map[string]map[string]float64{}
	datasetWithUniqueRatings :=  map[string]map[string]float64{}

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
		}else{
			for key := range otherRatings {

				userRatedItems = append(userRatedItems, key)
			}
		}
	}

	userItemDataset := assets.Data{ EqualUserItemRatings: &sameRatingsAsUser, UniqueUserItemRatings: &datasetWithUniqueRatings, AllUserItemRatings: uI.dataset}
	return userItemDataset, userRatedItems
}

// neighbours method to calculate nearestNeighbours
func (neighbours *nearestNeighbour) calculate(){
	//todo create algorithm to decide the unique items needed (Mean) for better results
	//todo check if user item has been rated multiple times by other users
	//todo check items user compaired with other users\
	//variables
	sim := neighbours.similarity.Dataset
	uniqueUserItems := *neighbours.data.UniqueUserItemRatings
	totalNeighbours:= 3
	sortedMap := sortedmap.New(totalNeighbours, asc.Float64)
	var nearestNeighbours map[string]float64
	nearestNeighbours = map[string]float64{}

	// walk true every similarity in the dataset
	for key, value := range sim{
		// only adds user when similarity is higher then the threshold
		if value > neighbours.threshold {
			// only adds user if it has more than 1 unique item rated
			if len(uniqueUserItems[key]) > 1 {
				nearestNeighbours[key] = value
				sortedMap.Insert(key, value)
			}
		}
		// break after finding totalNeighbours
		if len(nearestNeighbours) == totalNeighbours{
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
func (neighbours *nearestNeighbour) predictUniqueItemRatings()  {
	//variables
	ItemsPearsonRanked :=  map[string]float64{}
	listTimesOfAllItems := map[string]float64{}
	dataset := *neighbours.nearest

	//count the similarity for every unique item.
	for userID, algorithmDistance := range neighbours.similarity.Dataset {
		for key, value := range dataset[userID] {
			// add items to the list to create a specific search later
			predictedRatingForItem := value * algorithmDistance
			ItemsPearsonRanked[key] += predictedRatingForItem
			listTimesOfAllItems[key] += 1.0
		}
	}

	sortedMap := sortedmap.New(len(ItemsPearsonRanked), asc.Float64)

	// divide the sum of the item similarity to predict the users items rating
	for key := range ItemsPearsonRanked  {
		totalRating := ItemsPearsonRanked[key]
		timesRated := listTimesOfAllItems[key]
		ItemsPearsonRanked[key] = totalRating / timesRated
		sortedMap.Insert(key, totalRating / timesRated)
	}


	iterCh, err := sortedMap.BoundedIterCh(true, 0.0, 10.0)
	if err != nil {
		fmt.Println("error while sorting list", err)
	}

	fmt.Println("Predicted Items result:")
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
		fmt.Println(integer + " : Item: " + key + " Rating: " + value + ", ")
		if i == 10{
			break
		}
	}
}


func main() {
	//variables
	userID := "7"
	threshold := 0.8
	var list []assets.NewUserItemDataSet
	userRatings := assets.ReadDataset("files/item-item.txt")
	//todo uncomment if you wanna read the MovieDataSet
	//userRatings := assets.ReadMovieDataSet("files/movieLens100KUserItems.data")

	// PART 1
	userSeven := UserItem{userID, &userRatings}

	// PART 2
	// get the result off A algorithm on the dataset
	euclideanResult := userSeven.euclidean()
	pearsonResult := userSeven.pearson()
	cosineResult := userSeven.cosine()

	// create a list off all the results
	list = append(list, pearsonResult)
	list = append(list, euclideanResult)
	list = append(list, cosineResult)

	// print result off the algorithms
	assets.PrintMultipleAlgorithms(list,"The distance from user " + userID + " compared with the other users:")

	// PART 3
	// find similar and unique ratings
	equalAndUniqueUserItemRatings, userRatedItems := userSeven.findUsersWithMoreUniqueRatings()
	// print result off the dataset with the same and different ratings.
	assets.PrintsSimilarAndDifferentItems(equalAndUniqueUserItemRatings, "See the same and unique ratings for each user compared with user " + userID + ".")

	// Part 4
	// Nearest Neighbour
	nearestNeighbour := nearestNeighbour{userID, equalAndUniqueUserItemRatings, pearsonResult,threshold, nearestNeighbour{}.nearest, userRatedItems}
	nearestNeighbour.calculate()

	// Part 5
	nearestNeighbour.predictUniqueItemRatings()

	//todo extra create flexible threshold
}
