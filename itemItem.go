package main

import (
	"fmt"
	"hro.projects/INFDAT01-2NEW/algorithms"
	"hro.projects/INFDAT01-2NEW/assets"
	"sort"
)

type ItemItem struct {
	userID   string
	dataset  *map[string]map[string]float64
	average  *map[string]float64
	itemsIDs *[]string
}

func (iI ItemItem) createMatrix() {
	average := *iI.average
	itemIDs := *iI.itemsIDs
	for key, value := range *iI.dataset {
		counter := 0.0
		for key2, value2 := range value {
			if assets.ItemContains(itemIDs, key2) {
				itemIDs = append(itemIDs, key2)
			}
			average[key] += value2
			counter++
		}
		average[key] = average[key] / counter
	}
	sort.Strings(itemIDs)
	iI.average = &average
	*iI.itemsIDs = itemIDs
}

func (iI ItemItem) cosineAdjustmentFormula() {
	usersData := *iI.dataset
	var test []string
	test = *iI.itemsIDs
	var demo []string
	demo = *iI.itemsIDs

	vertex := algorithms.CosineAdjustment(usersData, *iI.average, test)
	assets.PrintVertexTable(vertex, demo)

}

func main() {
	//variables
	userID := "7"
	userRatings := assets.ReadDataset("files/item-item.txt")
	//userRatings := assets.ReadMovieDataSet("files/movieLens100KUserItems.data")

	//todo : Opstellen van een tabel waarbij alle artikelen met elkaar worden vergeleken
	// PART 1
	userSeven := ItemItem{userID, &userRatings, &map[string]float64{}, &[]string{}}
	userSeven.createMatrix()
	assets.PrintItemTable(userSeven.dataset, userSeven.average, *userSeven.itemsIDs, "avarage table")
	userSeven.cosineAdjustmentFormula()
	fmt.Print("test")
}
