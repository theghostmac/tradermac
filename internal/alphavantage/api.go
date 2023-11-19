package alphavantage

import "net/http"

// APIClient is an AlphaVantage API Client.
type APIClient struct {
	config     *Configuration
	httpClient *http.Client
}

// NewAPIClient instantiates an APIClient type for access to AlphaVantage.
func NewAPIClient(config *Configuration) *APIClient {
	return &APIClient{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// CallAlphaVantageAPI executes HTTP requests to AlphaVantageAPI using APIRequest. The response is a raw JSON.
func (ac *APIClient) CallAlphaVantageAPI(apiReq *APIRequest) (string, error) {

}
