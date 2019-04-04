package assets

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Apriori struct {
	OrderDate     time.Time
	TransactionID int
	ItemName      string
}

func ReadDataset(fileName string) (dataset map[string]map[string]float64, err error) {
	// Get file and collect the data from file
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error when trying to open the file.")
		log.Fatal(err)
	}

	// Split all content in lines by new line
	lines := strings.Split(string(content), "\n")
	if len(lines) == 1 {
		return dataset, fmt.Errorf("Database is nog leeg er kan geen recomendatie gedaan worden.")
	}
	// Create hashTable based on user int and items
	dataset = map[string]map[string]float64{}

	for i := 0; i < len(lines)-1; i++ {
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
	return dataset, nil
}

func ReadMovieDataSet(fileName string) (dataset map[string]map[string]float64, err error) {
	// Get file and collect the data from file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error loading MovieLens file: ", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	if len(scanner.Text()) == 1 {
		return dataset, fmt.Errorf("Database is nog leeg er kan geen recomendatie gedaan worden.")
	}

	// Create hashTable based on user int and items
	dataset = map[string]map[string]float64{}

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

	return dataset, nil
}

// needed to convert the Apriori dataset to the right dataset format.
func ReadAprioriDataSet(fileName string) map[int]Apriori {
	// Get file and collect the data from file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error loading MovieLens file: ", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Create hashTable based on user int and items
	dataset := map[int]Apriori{}
	integer := 0

	// use scanner to detect spaces and new lines.
	for scanner.Scan() {
		// Get id, itemId & rating from line comma separated
		lines := strings.Split(scanner.Text(), "/n")
		for i := 0; i < len(lines); i++ {
			// Get id, itemId & rating from line comma separated
			line := strings.Split(string(lines[i]), ",")

			transactionID, _ := strconv.Atoi(line[2])

			location, _ := time.LoadLocation("Europe/Amsterdam")
			str, err := time.Parse("2006-01-02 15:04:05", line[0]+" "+line[1])

			if err != nil {
				fmt.Println(err)
			}

			a := Apriori{str.In(location), transactionID, line[3]}

			// Add new map for the user in the map with itemID(string) and ItemRating(float64)
			dataset[integer] = a
			integer++
		}
	}

	fmt.Println(dataset[0], dataset[0].TransactionID)
	return dataset
}

func ReadItemInformation(fileName string) map[string]string {
	// Get file and collect the data from file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error loading item data file: ", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Create hashTable based on user int and items
	dataset := map[string]string{}
	//integer := 0

	// use scanner to detect spaces and new lines.
	for scanner.Scan() {
		// Get id, itemId & rating from line comma separated
		lines := strings.Split(scanner.Text(), "/n")
		//for i := 0; i < len(lines); i++ {
		for _, value := range lines {
			// Get id, itemId & rating from line comma separated
			line := strings.Split(string(value), "|")
			if err != nil {
				fmt.Println(err)
			}

			dataset[line[0]] = line[1]
		}
	}

	return dataset
}
