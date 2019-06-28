package assets

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"math"
	"os"
	"strconv"
)

type NewUserItemDataSet struct {
	UserID        string
	Dataset       map[string]float64
	AlgorithmName string
}

type ItemItemDataSet struct {
	Dataset       map[string]map[string]float64
	Averages      map[string]float64
	AlgorithmName string
}

type Data struct {
	EqualUserItemRatings  *map[string]map[string]float64
	UniqueUserItemRatings *map[string]map[string]float64
	AllUserItemRatings    *map[string]map[string]float64
}

func PrintMultipleAlgorithms(allAlgorithms []NewUserItemDataSet, description string) {
	totalUsers := len(allAlgorithms[0].Dataset)
	user, _ := strconv.Atoi(allAlgorithms[0].UserID)

	var tableData [][]string
	var tableHeaders []string
	tableHeaders = append(tableHeaders, "Users")

	for i := 1; i < totalUsers+1; i++ {
		var resultFromUser []string
		resultFromUser = append(resultFromUser, strconv.Itoa(i))
		if i == user {
			totalUsers += 1
			fmt.Println("user " + strconv.Itoa(i) + ":  -")
			resultFromUser = append(resultFromUser, "-")
		} else {
			// return the result out of the dataset convert it to a string and add it to the array of strings in resultFromUser
			for iteration, allAlgorithmResult2 := range allAlgorithms {
				if iteration <= len(allAlgorithms) && i == 1 {
					tableHeaders = append(tableHeaders, allAlgorithmResult2.AlgorithmName)
				}
				currentUser := strconv.Itoa(i)
				resultFromUser = append(resultFromUser, fmt.Sprintf("%.15f", allAlgorithmResult2.Dataset[currentUser]))
			}

			tableData = append(tableData, resultFromUser)
		}
		if i == 10 {
			break
		}
	}

	// create Ascii table
	fmt.Println(description)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders) // set the headers for the table
	table.SetBorder(true)
	table.AppendBulk(tableData) // Append the user data to the table
	table.Render()
}

func PrintsSimilarAndDifferentItems(data Data, description string) {
	// variables
	var tableData [][]string
	tableHeaders := []string{"Users", "Same ratings", "Unique ratings"}
	equalRatings := *data.EqualUserItemRatings
	totalUsers := len(equalRatings)
	uniqueRatings := *data.UniqueUserItemRatings

	for i := 1; i < totalUsers+1; i++ {
		// variables
		var stringSameRatings string
		var stringUniqueRatings string
		userID := strconv.Itoa(i)

		var resultFromUser []string
		resultFromUser = append(resultFromUser, strconv.Itoa(i))

		// add all same user ratings to the string
		for key, value := range equalRatings[userID] {
			tempString := fmt.Sprintf("%v", value)
			stringSameRatings += "(" + key + ":" + tempString + ") "
		}

		// add all unique user ratings to the string
		for key, value := range uniqueRatings[userID] {
			tempString := fmt.Sprintf("%v", value)
			stringUniqueRatings += "(" + key + ":" + tempString + ") "
		}

		// add  string results from user to table string
		resultFromUser = append(resultFromUser, stringSameRatings)
		resultFromUser = append(resultFromUser, stringUniqueRatings)
		tableData = append(tableData, resultFromUser)
		if i == 6 {
			break
		}
	}

	// print the description
	fmt.Println("\n" + description)
	// create Ascii table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders) // set the headers for the table
	table.SetBorder(true)
	table.AppendBulk(tableData) // Append the user data to the table
	table.Render()
}

// Contains tells if item contains the same item in the comparedItem list.
func Contains(items map[string]float64, comparedItem map[string]float64) bool {
	for i := range items {
		if _, ok := comparedItem[i]; ok {
			return true
		}
	}
	return false
}

// Contains tells if item contains the same item in the comparedItem list.
func ContainsString(items []string, comparedItem string) bool {
	for _, val := range items {
		if val == comparedItem {
			return true
		}
	}
	return false
}

// Find tells if value is in the array of strings
func Find(items []string, comparedItem string) bool {
	for _, value := range items {
		if comparedItem == value {
			return true
		}
	}
	return false
}

// Create new user for item prediction
func CreateNewUser(Id string) (userRatings map[string]map[string]float64, err error) {
	userRatings = make(map[string]map[string]float64)
	userRatings[Id] = map[string]float64{}
	userRatings[Id]["104"] = 3.0
	userRatings[Id]["106"] = 5.0
	userRatings[Id]["107"] = 4.0
	userRatings[Id]["109"] = 1.0

	return userRatings, nil
}

// Normalize ratings
func NormalizeUserRatings(averagesRatingBetweenItems map[string]map[float64]int, ratings map[string]float64, userLowestRating float64, userHighestRating float64) (normalizedRating float64, err error) {
	upper := 0.0
	lower := 0.0
	simItems := 0.0

	// go over every user rated item
	for k, v := range ratings {
		// get similarity between items
		for similarityBetweenItems, _ := range averagesRatingBetweenItems[k] {
			simItems = similarityBetweenItems
			break
		}
		upper += normaliseValue(v, userLowestRating, userHighestRating) * simItems
		lower += math.Abs(simItems)
	}

	// return the deviation between upper and lower to get the normalized value for the user.
	return upper / lower, nil
}

// Normalise value from 0 - 5 to -1 and 1.
func normaliseValue(value float64, userLowestRating float64, userHighestRating float64) float64 {
	return 2*((value-userLowestRating)/(userHighestRating-userLowestRating)) - 1
}

// Create value between 0 - 5 from normalised value between -1 and 1
func DenormalizeValue(normalisedValue float64, userLowestRating float64, userHighestRating float64) float64 {
	return ((normalisedValue+1)/2)*(userHighestRating-userLowestRating) + 1
}

// Get the lowest and highest rated item from the user
func GetUserLowestAndHighestValue(userRatings map[string]float64) (low float64, high float64) {
	low = 5.0
	high = 0.0

	for _, v := range userRatings {
		if v < low {
			low = v
		}
		if v > high {
			high = v
		}
	}

	return low, high
}
