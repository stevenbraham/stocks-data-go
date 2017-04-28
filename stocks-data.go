package main

import (
	"fmt"
	"os"
	"stocks-data/api"
)

func main() {
	if len(os.Args) == 2 {
		symbol := os.Args[1]
		company := api.Lookup(symbol)
		fmt.Println(company.Name)
		fmt.Println(company.Exchange)
		fmt.Println(company.StockPrice)
	} else {
		fmt.Println("Usage: stocks-data [symbol]")
	}
}
