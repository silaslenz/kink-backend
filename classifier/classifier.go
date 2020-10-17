package classifier

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/navossoc/bayesian"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Trains based on csv with two columns, text and category
func Train() *bayesian.Classifier {
	categories := GetCategories()

	classifier := bayesian.NewClassifier(categories...)
	records := readCsvFile("/home/silenz/Projects/kink/category_training_data.csv", ';')
	for _, record := range records[1:] {
		title := record[0]
		category := record[1]
		classifier.Learn(strings.Fields(title), bayesian.Class(category))

	}
	return classifier
}

func readCsvFile(filePath string, separator rune) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = separator
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
func GetCategories() []bayesian.Class {
	categoriesFile, err := os.Open("classifier/categories.json")
	if err != nil {
		fmt.Println(err)
	}
	categoriesJson, err := ioutil.ReadAll(categoriesFile)
	if err != nil {
		fmt.Print(err)
	}

	var categories []bayesian.Class
	_ = json.Unmarshal(categoriesJson, &categories)
	return categories
}
