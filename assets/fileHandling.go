package assets

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type item struct{
	itemID string
	rating float64
}

func ReadDataset(fileName string) map[string]map[string]float64 {
	// Get file and collect the data from file
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error when trying to open the file.")
		log.Fatal(err)
	}

	// Split all content in lines by new line
	lines := strings.Split(string(content), "\n")

	// Create hashTable based on user int and items
	dataset :=  map[string]map[string]float64{}

	for i := 0; i < len(lines) - 1;  i++{
		// Get id, itemId & rating from line comma separated
		line := strings.Split(string(lines[i]), ",")

		userID := line[0]
		itemID := line[1]
		itemRating, _ := strconv.ParseFloat(line[2], 64)

		// Add new map for the user in the map with itemID(string) and ItemRating(float64)
		if dataset[userID] == nil {
			dataset[userID] = map[string]float64{}
		}
		dataset[userID][itemID] = itemRating
	}
	return dataset
}

func ReadMovieDataSet(fileName string) map[string]map[string]float64 {
	// Get file and collect the data from file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error loading MovieLens file: ", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Create hashTable based on user int and items
	dataset :=  map[string]map[string]float64{}

	// use scanner to detect spaces and new lines.
	for scanner.Scan() {
		// Get id, itemId & rating from line comma separated
		line := strings.Split(scanner.Text(), "\t")
		//s := ss[len(ss)-1]
		userID := line[0]
		itemID := line[1]
		itemRating, _ := strconv.ParseFloat(line[2], 64)

		// Add new map for the user in the map with itemID(string) and ItemRating(float64)
		if dataset[userID] == nil {
			dataset[userID] = map[string]float64{}
		}
		dataset[userID][itemID] = itemRating
	}

	return dataset
}
