package dataparsers

import (
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
		_, likely, _ := trainedClassifier.LogScores(strings.Fields(title))
		transactions = append(transactions, models.Transaction{
			Id:       i,
			Date:     record[0],
			Title:    title,
			Amount:   amount,
			Balance:  balance,
			Currency: "SEK",
			Category: string(categories[likely]),
		})
	}
	return transactions
}
