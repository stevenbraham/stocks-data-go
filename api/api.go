package api

import (
	"encoding/json"
	"io"
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

	var apiData []map[string]string
	json.NewDecoder(DoApiCall(COMPANY_LOOKUP, stockSymbol)).Decode(&apiData)
	if len(apiData) == 0 {
		panic("API Error")
	}
	company := models.Company{
		Name:        apiData[0]["Name"],
		Exchange:    apiData[0]["Exchange"],
		StockSymbol: stockSymbol,
	}
	company.StockPrice = StockPrice(stockSymbol)
	return company
}

func StockPrice(stockSymbol string) float32 {
	var apiData map[string]float32
	json.NewDecoder(DoApiCall(QUOTE_LOOKUP, stockSymbol)).Decode(&apiData)
	return apiData["LastPrice"]
}

func DoApiCall(method ApiMethod, stockSymbol string) io.Reader {
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

	return response.Body
}
