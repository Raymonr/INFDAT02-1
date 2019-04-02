package assets

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"hro.projects/INFDAT01-2NEW/algorithms"
	"os"
	"reflect"
	"strconv"
)

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

func PrintItemTable(dataset *map[string]map[string]float64, average *map[string]float64, itemIDs []string, description string) {
	// variables
	var tableData [][]string
	tableHeaders := []string{"Users", "Avarage"}

	// add itemID's to header
	for i := 0; i < len(itemIDs); i++ {
		tableHeaders = append(tableHeaders, itemIDs[i])
	}

	// go true every user in dataset
	for key, value := range *dataset {
		var resultFromUser []string
		resultFromUser = append(resultFromUser, key)

		// needed to add average off user
		for key2, value2 := range *average {
			if key == key2 {
				resultFromUser = append(resultFromUser, fmt.Sprintf("%.2f", value2))
			}
		}
		for i := 0; i < len(itemIDs); i++ {
			temp := itemIDs[i]
			if val, ok := value[temp]; ok {
				resultFromUser = append(resultFromUser, fmt.Sprintf("%.2f", val))
			} else {
				resultFromUser = append(resultFromUser, "-")
			}
		}

		tableData = append(tableData, resultFromUser)
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
	fmt.Println("deze data", dataset)

	// variables
	var tableData [][]string
	tableHeaders := []string{"item"}

	// add itemID's to header
	for i := 1; i < len(itemIDs); i++ {
		tableHeaders = append(tableHeaders, itemIDs[i])
	}

	// go true every user in dataset
	for key, vertex := range dataset {
		var resultFromUser []string
		total := strconv.Itoa(getFieldInteger(vertex, "totalRatings"))
		cosine := getFieldFloat(vertex, "cosine")
		if key == 0 {
			firstItem := getFieldString(vertex, "firstItem")
			resultFromUser = append(resultFromUser, firstItem, fmt.Sprintf("%.4f", cosine)+" ("+total+")")
		} else {
			resultFromUser = append(resultFromUser, fmt.Sprintf("%.4f", cosine)+" ("+total+")")
		}

		tableData = append(tableData, resultFromUser)
	}

	// create Ascii table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders) // set the headers for the table
	table.SetBorder(true)
	table.AppendBulk(tableData) // Append the user data to the table
	table.Render()
	fmt.Println("einde")
}
