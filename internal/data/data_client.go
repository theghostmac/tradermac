package data

import (
	"fmt"
	"github.com/theghostmac/tradermac/internal/alphavantage"
	"os"
	"time"
)

// currentStockPrice  is fetched from API
// strikePrice is based on user input or contract specification.
// volatility index is calculated.
// riskFreeRate is fetched from API.
// timeToExpiration is user input or contract specification.

type AlphaVantageClient struct {
	client *alphavantage.APIClient
}

// NewAlphaVantageClient instantiates an AlphaVantage APIClient.
func NewAlphaVantageClient() (*AlphaVantageClient, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	client := alphavantage.NewAPIClient(config)
	return &AlphaVantageClient{client: client}, nil
}

func (avClient *AlphaVantageClient) CallAlphaVantageAPI(apiReq *alphavantage.APIRequest) (string, error) {
	return avClient.client.CallAlphaVantageAPI(apiReq)
}

func NewConfig() (configuration *alphavantage.Configuration, err error) {
	apiUrl := os.Getenv("ALPHAVANTAGE_API_URL")
	keyName := os.Getenv("ALPHAVANTAGE_KEY_NAME")
	apiKeyValue := os.Getenv("ALPHAVANTAGE_MARKET_API_KEY")

	if apiUrl == "" {
		return nil, fmt.Errorf("missing %s", "ALPHAVANTAGE_API_URL")
	}

	if keyName == "" {
		return nil, fmt.Errorf("missing %s", "ALPHAVANTAGE_KEY_NAME")
	}

	if apiKeyValue == "" {
		return nil, fmt.Errorf("missing %s", "ALPHAVANTAGE_MARKET_API_KEY")
	}

	config := alphavantage.NewConfiguration(apiUrl)
	config.AddKey(keyName, apiKeyValue)
	config.CustomTimeout(5 * time.Second)

	return config, nil
}
