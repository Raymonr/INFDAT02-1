package assets

import (
	"fmt"
	"github.com/cydev/zero"
	"time"
)

type transactions struct {
	orderDate     time.Time
	transactionID int
	itemNames     *[]string
}

func (transaction transactions) AddItem(item string) {
	*transaction.itemNames = append(*transaction.itemNames, item)
}

type Product map[string]int
type Candicates struct {
	itemSet      [2]string
	supportCount int
}

func FilterByTransaction(dataset map[int]Apriori) map[int]transactions {

	transactionList := map[int]transactions{}

	for _, value := range dataset {
		//place := transactionList[value.transactionID]

		// Add new map for the user in the map with itemID(string) and ItemRating(float64)
		if zero.IsZero(transactionList[value.TransactionID]) {
			newItem := []string{value.ItemName}
			transactionList[value.TransactionID] = transactions{value.OrderDate, value.TransactionID, &newItem}
		} else {
			place := transactionList[value.TransactionID]
			place.AddItem(value.ItemName)
		}
	}
	return transactionList
}

func CountNumberOffItems(dataset map[int]transactions) Product {
	product := Product{}
	for _, value := range dataset {
		for _, value2 := range *value.itemNames {
			fmt.Print(value)
			product[value2] += 1
		}
	}

	return product
}

func Threshold(products Product, threshold int) Product {
	newProduct := Product{}

	for key, value := range products {
		if value > threshold {
			newProduct[key] = value
		}
	}

	return newProduct
}

func Candidates(products Product) Candicates {
	newCandidates := Candicates{}
	for key, value := range products {
		fmt.Print(key, value, newCandidates)
	}

	return newCandidates
}
