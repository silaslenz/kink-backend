package dataparsers

import (
	"fmt"
	"kink_api/classifier"
	"kink_api/models"
	"strings"
)

func parseAmounts(amountString string) string {
	amountFields := strings.Fields(amountString)
	amountFields = amountFields[:len(amountFields)-1]
	amount := strings.ReplaceAll(strings.Join(amountFields, ""), ",", ".")
	return amount
}

func ReadIcaBanken(filePath string) []models.Transaction {
	trainedClassifier := classifier.Train()
	categories := classifier.GetCategories()
	var transactions []models.Transaction
	records := readCsvFile(filePath, ';')
	for i, record := range records[1:] {
		title := record[1]
		amount := parseAmounts(record[4])
		balance := parseAmounts(record[5])
		scores, likely, _ := trainedClassifier.ProbScores(strings.Fields(title))
		max := max(scores)

		category := "-"
		if max > 0.5 {
			category = string(categories[likely])
		}
		fmt.Println(max, title, category)

		transactions = append(transactions, models.Transaction{
			Id:       i,
			Date:     record[0],
			Title:    title,
			Amount:   amount,
			Balance:  balance,
			Currency: "SEK",
			Category: category,
		})
	}
	return transactions
}

func max(array []float64) float64 {
	max := array[0]
	for i := 1; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	return max
}
