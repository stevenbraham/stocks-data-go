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
	apiData := DoApiCall(COMPANY_LOOKUP, stockSymbol)
	if len(apiData) == 0 {
		panic("API Error")
	}
	company := models.Company{
		Name:        apiData[0]["Name"],
		Exchange:    apiData[0]["Exchange"],
		StockSymbol: stockSymbol,
	}
	return company
}

func DoApiCall(method ApiMethod, stockSymbol string) []map[string]string {
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
	var jsonData []map[string]string
	json.NewDecoder(response.Body).Decode(&jsonData)
	return jsonData
}
