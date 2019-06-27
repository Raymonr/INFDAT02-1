package algorithms

type Vertex struct {
	firstItem, secondItem string
	cosine                float64
	totalRatings          int
}

//func CosineAdjustment(dataset map[string]map[string]float64, average map[string]float64, itemIDs []string) []Vertex {
//	var vertex []Vertex
//	newList := itemIDs
//	// Calculates the distance between the same items and adds similarity to see the difference in total ratings
//	for i := 0; i < len(itemIDs); i++ {
//		// go true the whole list
//		var resultItemItem []Vertex
//		var itemName string
//		for j := 0; j < len(newList); j++ {
//			// check if the iterator is not at the end of the list else change it to the first.
//			if len(newList)-1 == j {
//				itemName = itemIDs[0]
//			} else {
//				itemName = itemIDs[j+1]
//			}
//			// check if items arn't the same then skip
//			if itemIDs[i] != itemName {
//				var A, B, C, D float64
//				counter := 0
//
//				// walk true every user to see if they rated the items
//				for key, value := range dataset {
//					var userAverage float64
//					// get average from user current
//
//					for keyAverage, valueAverage := range average {
//						if key == keyAverage {
//							userAverage = valueAverage
//							break
//						}
//					}
//					// check if first item and second item are rated by the user
//					if firstItem, ok := value[itemIDs[i]]; ok {
//						if secondItem, ok := value[itemName]; ok {
//							first := firstItem - userAverage
//							second := secondItem - userAverage
//							A += (first) * (second)
//							B += math.Pow(first, 2)
//							C += math.Pow(second, 2)
//							counter++
//						}
//					}
//				}
//				D = A / (math.Sqrt(B) * math.Sqrt(C))
//				// append result of all users on the first and second item
//				resultItemItem = append(resultItemItem, Vertex{itemIDs[i], itemName, D, counter})
//			}
//		}
//		vertex = resultItemItem
//		if len(newList) != 2 {
//			//newList[i] = newList[len(newList)-1] // Copy last element to index i.
//			newList = newList[:len(newList)-1]
//		} else {
//			break
//		}
//	}
//
//	return vertex
//}
