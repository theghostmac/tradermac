package main

import (
	"fmt"
	"github.com/theghostmac/tradermac/internal/alphavantage"
	"github.com/theghostmac/tradermac/internal/data"
	"log"
	"os"
)

func main() {
	// Create a new AlphaVantageClient
	avClient, err := data.NewAlphaVantageClient()
	if err != nil {
		log.Fatalf("Error creating AlphaVantage client: %v", err)
	}

	// Example: Fetch global quote data
	symbol := "AAPL" // Replace with the desired stock symbol

	// Log API key for debugging
	apiKey := os.Getenv("ALPHAVANTAGE_MARKET_API_KEY")
	log.Printf("Using API Key: %s", apiKey)

	globalQuoteData, err := fetchGlobalQuote(avClient, symbol)
	if err != nil {
		log.Fatalf("Error fetching global quote data: %v", err)
	}
	fmt.Println("Global Quote Data:")
	fmt.Println(globalQuoteData)

	// Example: Fetch time series daily data
	timeSeriesDailyData, err := fetchTimeSeriesDaily(avClient, symbol)
	if err != nil {
		log.Fatalf("Error fetching time series daily data: %v", err)
	}
	fmt.Println("Time Series Daily Data:")
	fmt.Println(timeSeriesDailyData)
}

func fetchGlobalQuote(avClient *data.AlphaVantageClient, symbol string) (string, error) {
	apiReq := alphavantage.NewAPIRequest().
		SetAPIReqMethod("GET").
		SetAPIReqPath("query").
		SetAPIReqParam("function", "GLOBAL_QUOTE").
		SetAPIReqParam("symbol", symbol)

	response, err := avClient.CallAlphaVantageAPI(apiReq)
	if err != nil {
		return "", err
	}

	return response, nil
}

func fetchTimeSeriesDaily(avClient *data.AlphaVantageClient, symbol string) (string, error) {
	apiReq := alphavantage.NewAPIRequest().
		SetAPIReqMethod("GET").
		SetAPIReqPath("query").
		SetAPIReqParam("function", "TIME_SERIES_DAILY").
		SetAPIReqParam("symbol", symbol)

	response, err := avClient.CallAlphaVantageAPI(apiReq)
	if err != nil {
		return "", err
	}

	return response, nil
}
