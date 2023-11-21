package data

import (
	"encoding/json"
	"github.com/theghostmac/tradermac/internal/alphavantage"
	"log"
)

// MarketData is responsible for retrieving market data.
type MarketData struct {
	alphavantageClient *alphavantage.APIClient
}

// NewMarketData creates a new instance of MarketData.
func NewMarketData(alphavantageClient *alphavantage.APIClient) *MarketData {
	return &MarketData{alphavantageClient: alphavantageClient}
}

// GetCurrentStockPrice retrieves the current stock price for a given symbol.
func (md *MarketData) GetCurrentStockPrice(symbol string) (float64, error) {
	apiRequest := alphavantage.NewAPIRequest().
		SetAPIReqMethod("GET").
		SetAPIReqPath("/stock/quote").
		SetAPIReqParam("symbol", symbol)

	response, err := md.alphavantageClient.CallAlphaVantageAPI(apiRequest)
	if err != nil {
		log.Fatalf("Failed to call the alpha vantage api: %v", err)
		return 0, err
	}

	var quote struct {
		Price float64 `json:"price"`
	}

	err = json.Unmarshal([]byte(response), &quote)
	if err != nil {
		log.Fatalf("Failed to parse the response: %v", err)
		return 0, err
	}
	return quote.Price, nil
}
