package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"kink_api/models"
	"time"
)

func Sum(transactions []models.Transaction, from time.Time, to time.Time) decimal.Decimal {
	layout := "2006-01-02"
	sum := decimal.NewFromInt(0)
	println(from.String(), to.String())
	for _, transaction := range transactions {
		date, err := time.Parse(layout, transaction.Date)
		if err != nil {
			fmt.Println(err)
		}
		if dateBetween(date, from, to) {
			println(transaction.Date, transaction.Title)
			amount, err := decimal.NewFromString(transaction.Amount)
			if err != nil {
				fmt.Println(err)
			}
			sum = sum.Add(amount)
		}
	}
	return sum
}

func dateBetween(date time.Time, from time.Time, to time.Time) bool {
	return (date.After(from) || date.Equal(from)) && date.Before(to)
}
