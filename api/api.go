package api

import (
	"encoding/json"
	"net/http"
	"stocks-data/models"
	"time"
)

type ApiMethod int

const (
	COMPANY_LOOKUP ApiMethod = 1 + iota
	QUOTE_LOOKUP
)

func Lookup(stockSymbol string) models.Company {
	DoApiCall(COMPANY_LOOKUP, stockSymbol)
	return models.Company{
		Name:        "Google INC",
		StockPrice:  1,
		StockSymbol: stockSymbol,
		Exchange:    "NASDAQ",
	}
}

func DoApiCall(method ApiMethod, stockSymbol string) interface{} {
	//prepare url call
	apiUrl := "http://dev.markitondemand.com/MODApis/Api/v2/"
	//append appropiate arguments
	switch method {
	case COMPANY_LOOKUP:
		apiUrl += "Lookup/json?input=" + stockSymbol
	case QUOTE_LOOKUP:
		apiUrl += "Quote/json?symbol=" + stockSymbol
	default:
		panic("Invalid api method")
	}
	apiClient := http.Client{
		Timeout: time.Second * 3,
	}

	request, _ := http.NewRequest(http.MethodGet, apiUrl, nil)

	response, error := apiClient.Do(request)

	if error != nil {
		panic("API error")
	}

	//blank struct to store json data
	var i interface{}
	return json.NewDecoder(response.Body).Decode(&i)
}
