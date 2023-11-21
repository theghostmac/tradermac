# stuff

```json
{
  "Meta Data": {
    "1. Information": "Daily Prices (open, high, low, close) and Volumes",
    "2. Symbol": "IBM",
    "3. Last Refreshed": "2023-11-15",
    "4. Output Size": "Compact",
    "5. Time Zone": "US/Eastern"
  },
  "Time Series (Daily)": {
    "2023-11-15": {
      "1. open": "150.4000",
      "2. high": "153.2200",
      "3. low": "150.4000",
      "4. close": "152.5800",
      "5. volume": "4632519"
    }
  }
}
```

```go
// data_client.go

package data

import (
	"encoding/json"
	"github.com/yourusername/yourproject/internal/alphavantage"
)

// TimeSeriesData represents the complete time series data from AlphaVantage.
type TimeSeriesData struct {
	MetaData   map[string]string `json:"Meta Data"`
	TimeSeries map[string]struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	} `json:"Time Series (Daily)"`
}

// GetCurrentStockTimeSeries retrieves the complete time series data for a given symbol.
func (c *MarketDataClient) GetCurrentStockTimeSeries(symbol string) (*TimeSeriesData, error) {
	apiRequest := alphavantage.NewAPIRequest().
		SetAPIReqMethod("GET").
		SetAPIReqPath("/time-series").
		SetAPIReqParam("symbol", symbol)

	response, err := c.alphavantageClient.CallAlphaVantageAPI(apiRequest)
	if err != nil {
		return nil, err
	}

	var timeSeriesData TimeSeriesData

	err = json.Unmarshal([]byte(response), &timeSeriesData)
	if err != nil {
		return nil, err
	}

	return &timeSeriesData, nil
}
```

```go
// main.go

package main

import (
	"fmt"
	"github.com/yourusername/yourproject/internal/alphavantage"
	"github.com/yourusername/yourproject/internal/data"
	"os"
	"time"
)

func main() {
	// Initialize AlphaVantage client
	client, err := data.NewAlphaVantageClient()
	if err != nil {
		fmt.Println("Error creating AlphaVantage client:", err)
		return
	}

	// Initialize MarketDataClient
	marketDataClient := data.NewMarketDataClient(client)

	// Get time series data for a stock symbol (e.g., "AAPL")
	symbol := "AAPL"
	timeSeriesData, err := marketDataClient.GetCurrentStockTimeSeries(symbol)
	if err != nil {
		fmt.Println("Error getting time series data:", err)
		return
	}

	// Print some sample data (you can modify this based on your actual response structure)
	fmt.Printf("Symbol: %s\n", symbol)
	fmt.Printf("First Open Price: %s\n", timeSeriesData.TimeSeries["2023-11-20"].Open)
	fmt.Printf("First Close Price: %s\n", timeSeriesData.TimeSeries["2023-11-20"].Close)
}
```