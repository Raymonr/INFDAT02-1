package assets

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"hro.projects/INFDAT01-2NEW/algorithms"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
)

// function to check if value exist.
func ItemContains(strArray []string, find string) bool {
	if len(strArray) == 0 {
		return true
	}

	for _, value := range strArray {
		if value == find {
			return false
		}
	}
	return true
}

// function to print interface averages
func PrintItemItemAverages(allAlgorithms map[string]float64, description string) {
	// variables
	var tableData [][]string
	tableHeaders := []string{"users", "average"}

	// add itemID's to header
	for key, value := range allAlgorithms {
		test := fmt.Sprintf("%.2f", value)
		var temp []string
		temp = append(temp, key)
		temp = append(temp, test)
		tableData = append(tableData, temp)
	}

	fmt.Println(description)
	// create Ascii table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders) // set the headers for the table
	table.SetBorder(true)
	table.AppendBulk(tableData) // Append the user data to the table
	table.Render()
}

// function to print interface similarities
func PrintItemAlgorithmSimilarities(cosineAdjustedSimilarity *map[string]map[string]map[float64]int, itemIDs []string, description string) {
	var tableData [][]string
	tableHeaders := []string{"Items"}

	for i := 0; i < len(itemIDs); i++ {
		if i+1 == len(itemIDs) {
			break
		}
		tableHeaders = append(tableHeaders, itemIDs[i+1])
	}

	counter := 0
	// add Similarities to table
	for key, value := range *cosineAdjustedSimilarity {
		counter++
		lenID := len(itemIDs)
		// break when last column gets repeated
		if key == itemIDs[lenID-1] {
			continue
		}
		// create row string for table
		var temp []string
		temp = append(temp, key)

		// add spaces to array string for the right position of item
		positionKeyFromItemsID := sort.StringSlice(itemIDs).Search(key)
		for i := 0; i < positionKeyFromItemsID; i++ {
			temp = append(temp, "")
			counter++
		}

		// loop over all values connected with key
		for _, value2 := range value {
			// get similarity and totalUsers from Value2

			for k, v := range value2 {
				temp = append(temp, fmt.Sprintf("%.2f", k)+"("+strconv.Itoa(v)+")")
				break
			}
		}
		// add the array of strings tot the row.
		tableData = append(tableData, temp)
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

func getFieldString(e algorithms.Vertex, field string) string {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func getFieldFloat(e algorithms.Vertex, field string) float64 {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Float()
}

func getFieldInteger(e algorithms.Vertex, field string) int {
	r := reflect.ValueOf(e)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func PrintVertexTable(dataset []algorithms.Vertex, itemIDs []string) {
	// variables
	var tableData [][]string
	tableHeaders := []string{"item"}

	// add itemID's to header
	for i := 1; i < len(itemIDs); i++ {
		tableHeaders = append(tableHeaders, itemIDs[i])
	}

	// go true every user in dataset
	var resultFromUser []string
	for key, vertex := range dataset {
		total := strconv.Itoa(getFieldInteger(vertex, "totalRatings"))
		cosine := getFieldFloat(vertex, "cosine")
		if key == 0 {
			firstItem := getFieldString(vertex, "firstItem")
			resultFromUser = append(resultFromUser, firstItem)
		}
		resultFromUser = append(resultFromUser, fmt.Sprintf("%.4f", cosine)+" ("+total+")")
	}
	tableData = append(tableData, resultFromUser)

	fmt.Println("\n")
	// create Ascii table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders) // set the headers for the table
	table.SetBorder(true)
	table.AppendBulk(tableData) // Append the user data to the table
	table.Render()
	fmt.Println("einde")
}

// function which orders the map from low to high id's
func SortList(usersAverages map[string]float64) map[string]float64 {
	// To store the keys in slice in sorted order
	var keys []string
	for k := range usersAverages {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	newList := map[string]float64{}
	// To perform the operation you want
	for _, k := range keys {
		newList[k] = usersAverages[k]
	}

	return newList
}

// function to calculate adjusted cosine.
func ACS(adjustedCosineTable map[string]map[string]map[float64]int, dataset map[string]map[string]float64, itemIds []string, userAverages map[string]float64) (act map[string]map[string]map[float64]int, err error) {
	// loop over all possible items
	for k, v := range adjustedCosineTable {
		// loop over all combinations from the item above
		for key := range v {
			var counter = 0
			var upper = 0.0
			var lower = 0.0
			var userItemAPow = 0.0
			var userItemBPow = 0.0
			var lowA = 0.0
			var lowB = 0.0
			result := make(map[float64]int)

			// Loop over all users
			for userKey, usersVal := range dataset {
				userAverage := userAverages[userKey]

				//check if user has rated both items
				if _, ok := usersVal[k]; ok {
					if _, ok := usersVal[key]; ok {
						//ACS formula calculate user item with other user item
						userItemAPow += math.Pow(usersVal[k]-userAverage, 2)
						userItemBPow += math.Pow(usersVal[key]-userAverage, 2)
						upper += (usersVal[k] - userAverage) * (usersVal[key] - userAverage)
						counter++
					}
				}
			}

			// Calculate total user similarity between items
			lowA = math.Sqrt(userItemAPow)
			lowB = math.Sqrt(userItemBPow)
			lower = lowA * lowB

			// Count upper and lower part of the ACS and divide bij total users for the similarity
			total := upper / lower

			// Add result of all the similarity between the items rated by each user.
			result[total] = counter
			adjustedCosineTable[k][key] = result
		}
	}
	return adjustedCosineTable, nil
}

// function to print one slope deviations
func printDeviationsOneSlope(deviations map[string]map[string]dev, items []string) {
	fmt.Println("dev", deviations)

	// variables
	var tableData [][]string
	var temp []string
	tableHeaders := []string{"items"}

	for _, val := range items {
		tableHeaders = append(tableHeaders, val)

		temp = append(temp, val)

		for _, v := range items {
			// if item exist

			if _, ok := deviations[val][v]; ok {
				test := fmt.Sprintf("%.2f", deviations[val][v].difference) + " (" + strconv.Itoa(deviations[val][v].ratedAmount) + ")"
				temp = append(temp, test)
			} else {
				temp = append(temp, "")
			}

		}
		tableData = append(tableData, temp)
	}

	// add itemID's to header
	//for k, v := range deviations {
	//	tableHeaders = append(tableHeaders, k)
	//	var temp []string
	//
	//	// loop over all values in items
	//	for _, val := range items{
	//		first := true
	//		// if item exist
	//		if first {
	//			temp = append(temp, k)
	//			first = false
	//		}
	//
	//		if _, ok := v[val]; ok {
	//			test := fmt.Sprintf("%.2f", v[val].difference) + " (" + strconv.Itoa(v[val].ratedAmount) + ")"
	//			temp = append(temp, test)
	//		}else {
	//			temp = append(temp, "")
	//		}
	//	}

	//for key, val := range v {
	//		test := fmt.Sprintf("%.2f", val.difference) + " (" + strconv.Itoa(val.ratedAmount) + ")"
	//		var temp []string
	//		temp = append(temp, key)
	//		temp = append(temp, test)
	//		tableData = append(tableData, temp)
	//}
	//tableData = append(tableData, temp)
	//fmt.Println("deze", k , v)

	//for key, value := range allAlgorithms {
	//	test := fmt.Sprintf("%.2f", value)
	//	var temp []string
	//	temp = append(temp, key)
	//	temp = append(temp, test)
	//	tableData = append(tableData, temp)
	//}

	// create Ascii table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders) // set the headers for the table
	table.SetBorder(true)
	table.AppendBulk(tableData) // Append the user data to the table
	table.Render()
}
