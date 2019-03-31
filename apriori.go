package main

import (
	"fmt"
	"hro.projects/INFDAT01-2NEW/assets"
)

func main() {
	// the used datasheet is from:
	// https://www.kaggle.com/sulmansarwar/transactions-from-a-bakery

	//variables
	threshold := 200
	//todo
	// read dataset
	//aprioriDataset
	aprioriDataset := assets.ReadAprioriDataSet("files/apriori-BreadBasket.csv")
	//testMap := map[int]assets.Apriori{}
	//str, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	//testMap[0] = assets.Apriori{OrderDate: str, TransactionID: 1, ItemName: "apple"}
	//testMap[1] = assets.Apriori{OrderDate: str, TransactionID: 1, ItemName: "egg"}
	//testMap[2] = assets.Apriori{OrderDate: str, TransactionID: 1, ItemName: "milk"}
	//testMap[3] = assets.Apriori{OrderDate: str, TransactionID: 2, ItemName: "egg"}
	//testMap[4] = assets.Apriori{OrderDate: str, TransactionID: 2, ItemName: "milk"}
	//testMap[5] = assets.Apriori{OrderDate: str, TransactionID: 1, ItemName: "apple"}
	//testMap[6] = assets.Apriori{OrderDate: str, TransactionID: 3, ItemName: "apple"}
	//testMap[7] = assets.Apriori{OrderDate: str, TransactionID: 4, ItemName: "apple"}
	//testMap[8] = assets.Apriori{OrderDate: str, TransactionID: 6, ItemName: "egg"}
	//testMap[9] = assets.Apriori{OrderDate: str, TransactionID: 4, ItemName: "carrot"}
	//testMap[10] = assets.Apriori{OrderDate: str, TransactionID: 6, ItemName: "carrot"}
	//aprioriDataset := testMap

	//
	C1 := assets.FilterByTransaction(aprioriDataset)
	countedNumberOffItems := assets.CountNumberOffItems(C1)
	// Candidate set support count with candidates above the Threshold
	L1 := assets.Threshold(countedNumberOffItems, threshold)
	// Generated candidates
	C2 := assets.Candidates(L1)
	fmt.Println("demo", aprioriDataset, C1, countedNumberOffItems, L1, C2)
	// Support
	// Confidence
	// Lift

}
