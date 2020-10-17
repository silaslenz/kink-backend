package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"kink_api/models"
	"time"
)

func Sum(transactions []models.Transaction, from time.Time, to time.Time) map[string]decimal.Decimal {
	layout := "2006-01-02"
	sums := map[string]decimal.Decimal{}
	for _, transaction := range transactions {
		date, err := time.Parse(layout, transaction.Date)
		if err != nil {
			fmt.Println(err)
		}
		if dateBetween(date, from, to) {
			amount, err := decimal.NewFromString(transaction.Amount)
			if err != nil {
				fmt.Println(err)
			}
			sums[transaction.Category] = sums[transaction.Category].Add(amount)
			sums["total"] = sums["total"].Add(amount)
		}
	}
	return sums
}

func dateBetween(date time.Time, from time.Time, to time.Time) bool {
	return (date.After(from) || date.Equal(from)) && date.Before(to)
}
