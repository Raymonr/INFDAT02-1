package assets

import (
	"github.com/cydev/zero"
	"time"
)

type Transactions struct {
	orderDate     time.Time
	transactionID int
	itemNames     *[]string
}

type AprioriData struct {
	Dataset       FrequentThreePairs
	ConfidenceSet map[ThreePairs]int
	Transactions  map[int]Transactions
}

type Pair struct {
	first, second interface{}
	count         int
}

type ThreePairs struct {
	first, second, third interface{}
	count                int
}

type Product map[string]int
type FrequentPairs []Pair
type FrequentThreePairs []ThreePairs

//add item to a transaction pointer
func (transaction Transactions) AddItem(item string) {
	*transaction.itemNames = append(*transaction.itemNames, item)
}

//
func FilterByTransaction(dataset map[int]Apriori) map[int]Transactions {
	transactionList := map[int]Transactions{}

	for _, value := range dataset {
		// Add new map for the user in the map with itemID(string) and ItemRating(float64)
		if zero.IsZero(transactionList[value.TransactionID]) {
			newItem := []string{value.ItemName}
			transactionList[value.TransactionID] = Transactions{value.OrderDate, value.TransactionID, &newItem}
		} else {
			place := transactionList[value.TransactionID]
			place.AddItem(value.ItemName)
		}
	}
	return transactionList
}

// count the number of items a item is in the dataset
func CountNumberOffItems(dataset map[int]Transactions) Product {
	product := Product{}
	for _, value := range dataset {
		for _, value2 := range *value.itemNames {
			product[value2] += 1
		}
	}

	return product
}

// check if the product value is above threshold
func Threshold(products Product, threshold int) Product {
	newProduct := Product{}

	for key, value := range products {
		if value > threshold {
			newProduct[key] = value
		}
	}

	return newProduct
}

// returns true when pair doesn't exist in freqItems
func checkIfPairExist(freqItems FrequentPairs, pair Pair) bool {
	if len(freqItems) == 0 {
		return true
	}
	otherPair := Pair{pair.second, pair.first, 0}
	for _, value := range freqItems {
		if value == otherPair || value == pair {
			return false
		}
	}

	return true
}

// check for Candidates returns a pair of unique Candidates
func Candidates(products Product) FrequentPairs {
	freqItems := FrequentPairs{}
	for key := range products {
		for key2 := range products {
			if key != key2 {
				newPair := Pair{key, key2, 0}
				if checkIfPairExist(freqItems, newPair) {
					freqItems = append(freqItems, newPair)
				}
			}
		}
	}

	return freqItems
}

// this function checks if string contains a pair and returns true if pair exist and false if not
func check(values []string, pair Pair) bool {
	first := false
	second := false
	for i := range values {
		if values[i] == pair.first {
			first = true
		} else if values[i] == pair.second {
			second = true
		}
		if first == true && second == true {
			break
		}
	}

	if first == true && second == true {
		return true
	} else {
		return false
	}
}

// this function checks if string contains a pair and returns true if pair exist and false if not
func CheckThree(values []string, pair ThreePairs) bool {
	first := false
	second := false
	third := false
	for i := range values {
		if values[i] == pair.first {
			first = true
		} else if values[i] == pair.second {
			second = true
		} else if values[i] == pair.third {
			third = true
		}
		if first == true && second == true && third == true {
			break
		}
	}

	if first == true && second == true && third == true {
		return true
	} else {
		return false
	}
}

// this function counts the the pairs that where in Transactions
func CountTransactionsFromPairs(frequentPairs *FrequentPairs, products map[int]Transactions, threshold int) FrequentPairs {
	newFrequentPairs := FrequentPairs{}
	for key, value2 := range *frequentPairs {
		for _, value := range products {
			temp := *value.itemNames
			if check(temp, value2) {
				demo := *frequentPairs
				demo[key].count += 1
			}
		}
	}

	for key, value := range *frequentPairs {
		temp := *frequentPairs
		if temp[key].count >= threshold {
			newFrequentPairs = append(newFrequentPairs, value)
		}
	}

	return newFrequentPairs
}

// check if three pairs already exist return true if not.
func checkIfThreePairsExist(firstThreePair ThreePairs, pairs FrequentThreePairs) bool {
	if len(pairs) == 0 {
		return true
	}
	tp := firstThreePair

	secondPair := ThreePairs{tp.third, tp.second, tp.first, 0}
	thirdPair := ThreePairs{tp.second, tp.third, tp.first, 0}
	fourthPair := ThreePairs{tp.first, tp.third, tp.second, 0}
	fifthPair := ThreePairs{tp.second, tp.first, tp.third, 0}
	sixthPair := ThreePairs{tp.third, tp.first, tp.second, 0}
	for _, value := range pairs {
		if value == tp || value == secondPair || value == thirdPair || value == fourthPair || value == fifthPair || value == sixthPair {
			return false
		}
	}

	return true
}

func FindThreePairs(frequentPairs *FrequentPairs) FrequentThreePairs {
	newThreePairs := FrequentThreePairs{}
	for key, value := range *frequentPairs {
		for key2, value2 := range *frequentPairs {
			// when it's the same key continue to the next
			if key == key2 {
				break
			}

			// if value is equal
			if value.first == value2.first {
				newThreePair := ThreePairs{value.first, value.second, value2.second, 0}
				if checkIfThreePairsExist(newThreePair, newThreePairs) {
					newThreePairs = append(newThreePairs, newThreePair)
				}
			} else if value.second == value2.first {
				newThreePair := ThreePairs{value.first, value.second, value2.second, 0}
				if checkIfThreePairsExist(newThreePair, newThreePairs) {
					newThreePairs = append(newThreePairs, newThreePair)
				}
			} else if value.first == value2.second {
				newThreePair := ThreePairs{value.first, value.second, value2.first, 0}
				if checkIfThreePairsExist(newThreePair, newThreePairs) {
					newThreePairs = append(newThreePairs, newThreePair)
				}
			} else if value.second == value2.second {
				newThreePair := ThreePairs{value.first, value.second, value2.first, 0}
				if checkIfThreePairsExist(newThreePair, newThreePairs) {
					newThreePairs = append(newThreePairs, newThreePair)
				}
			}
		}
	}

	return newThreePairs
}

func CountTransactionsFromThreePairs(frequentPairs FrequentThreePairs, products map[int]Transactions, threshold int) FrequentThreePairs {
	newFrequentPairs := FrequentThreePairs{}
	for key, value2 := range frequentPairs {
		for _, value := range products {
			temp := *value.itemNames
			if CheckThree(temp, value2) {
				demo := frequentPairs
				demo[key].count += 1
			}
		}
	}

	for key, value := range frequentPairs {
		temp := frequentPairs
		if temp[key].count >= threshold {
			newFrequentPairs = append(newFrequentPairs, value)
		}
	}

	return newFrequentPairs
}

func (apriori AprioriData) Confidence() {
	//Confidence(A->B)=Support_count(A∪B)/Support_count(A)
	for _, value := range apriori.Dataset {
		for _, value3 := range apriori.Transactions {
			items := *value3.itemNames
			val1 := false
			val2 := false
			val3 := false
			for i := 0; i < len(items); i++ {
				if items[i] == value.first {
					val1 = true
				} else if items[i] == value.second {
					val2 = true
				} else if items[i] == value.third {
					val3 = true
				}
				if val1 == true && val2 == true && val3 == true {
					break
				}
			}

			if val1 == true && val2 == true && val3 == true {
				apriori.ConfidenceSet[value] += 1
			}
		}
	}
}
