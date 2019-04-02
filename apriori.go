package main

import (
	"fmt"
	"hro.projects/INFDAT01-2NEW/assets"
	"time"
)

func main() {
	// the used datasheet is from:
	// https://www.geeksforgeeks.org/apriori-algorithm/
	// https://www.kaggle.com/sulmansarwar/transactions-from-a-bakery

	//variables
	threshold := 1
	//todo
	// read dataset
	//aprioriDataset
	//aprioriDataset := assets.ReadAprioriDataSet("files/apriori-BreadBasket.csv")
	testMap := map[int]assets.Apriori{}
	str, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	testMap[0] = assets.Apriori{OrderDate: str, TransactionID: 100, ItemName: "i1"}
	testMap[1] = assets.Apriori{OrderDate: str, TransactionID: 100, ItemName: "i2"}
	testMap[2] = assets.Apriori{OrderDate: str, TransactionID: 100, ItemName: "i5"}
	testMap[3] = assets.Apriori{OrderDate: str, TransactionID: 200, ItemName: "i2"}
	testMap[4] = assets.Apriori{OrderDate: str, TransactionID: 200, ItemName: "i4"}
	testMap[5] = assets.Apriori{OrderDate: str, TransactionID: 300, ItemName: "i2"}
	testMap[6] = assets.Apriori{OrderDate: str, TransactionID: 300, ItemName: "i3"}
	testMap[7] = assets.Apriori{OrderDate: str, TransactionID: 400, ItemName: "i1"}
	testMap[8] = assets.Apriori{OrderDate: str, TransactionID: 400, ItemName: "i2"}
	testMap[9] = assets.Apriori{OrderDate: str, TransactionID: 400, ItemName: "i4"}
	testMap[10] = assets.Apriori{OrderDate: str, TransactionID: 500, ItemName: "i1"}
	testMap[11] = assets.Apriori{OrderDate: str, TransactionID: 500, ItemName: "i3"}
	testMap[12] = assets.Apriori{OrderDate: str, TransactionID: 600, ItemName: "i2"}
	testMap[13] = assets.Apriori{OrderDate: str, TransactionID: 600, ItemName: "i3"}
	testMap[14] = assets.Apriori{OrderDate: str, TransactionID: 700, ItemName: "i1"}
	testMap[15] = assets.Apriori{OrderDate: str, TransactionID: 700, ItemName: "i3"}
	testMap[16] = assets.Apriori{OrderDate: str, TransactionID: 800, ItemName: "i1"}
	testMap[17] = assets.Apriori{OrderDate: str, TransactionID: 800, ItemName: "i2"}
	testMap[18] = assets.Apriori{OrderDate: str, TransactionID: 800, ItemName: "i3"}
	testMap[19] = assets.Apriori{OrderDate: str, TransactionID: 800, ItemName: "i5"}
	testMap[20] = assets.Apriori{OrderDate: str, TransactionID: 900, ItemName: "i1"}
	testMap[21] = assets.Apriori{OrderDate: str, TransactionID: 900, ItemName: "i2"}
	testMap[22] = assets.Apriori{OrderDate: str, TransactionID: 900, ItemName: "i3"}
	aprioriDataset := testMap

	//
	C1 := assets.FilterByTransaction(aprioriDataset)
	countedNumberOffItems := assets.CountNumberOffItems(C1)
	// Candidate set support count with candidates above the Threshold
	L1 := assets.Threshold(countedNumberOffItems, threshold)
	// Generated candidates
	C2 := assets.Candidates(L1)
	// Scan in transaction how much candidates are ordered and return the candidates above threshold
	L2 := assets.CountTransactionsFromPairs(&C2, C1, 2)
	// C3
	C3 := assets.FindThreePairs(&L2)
	L3 := assets.CountTransactionsFromThreePairs(C3, C1, 2)
	apr := assets.AprioriData{Dataset: L3, ConfidenceSet: map[assets.ThreePairs]int{}, Transactions: C1}
	apr.Confidence()
	fmt.Println("c3", L3, apr)
	// Support
	// Confidence
	// Lift

}
