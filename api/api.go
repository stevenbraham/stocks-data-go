package api

import (
	"stocks-data/models"
)

func Lookup(stockSymbol string) models.Company {
	return models.Company{
		Name:        "Google INC",
		StockPrice:  1,
		StockSymbol: stockSymbol,
		Exchange:    "NASDAQ",
	}
}
