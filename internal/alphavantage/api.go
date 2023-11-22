package alphavantage

import (
	"io"
	"log"
	"net/http"
)

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
	request, err := buildHttpRequest(apiReq, *ac.config)
	if err != nil {
		log.Printf("call: %s", err.Error())
		return "", err
	}

	log.Println(request.URL.String())
	log.Flags()

	httpClient := ac.httpClient
	response, err := httpClient.Do(request)
	if err != nil {
		log.Printf("call: %s", err.Error())
		return "", err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("call: %s", err.Error())
		return "", err
	}

	defer response.Body.Close()

	return string(body), nil
}

func buildHttpRequest(apiReq *APIRequest, config Configuration) (*http.Request, error) {
	endpoint := config.URL + apiReq.path
	req, err := http.NewRequest(apiReq.method, endpoint, nil)
	if err != nil {
		log.Println("Build HTTP Request: ", err)
		return nil, err
	}

	query := req.URL.Query()
	for key, value := range config.APIKeys {
		query.Add(key, value)
	}

	for param, value := range apiReq.params {
		query.Add(param, value)
	}

	req.URL.RawQuery = query.Encode()

	return req, nil
}

// func buildHttpRequest(apiReq *APIRequest, config Configuration) (*http.Request, error) {
// 	endpoint := config.URL + apiReq.path
// 	req, err := http.NewRequest(apiReq.method, endpoint, nil)
// 	if err != nil {
// 		log.Println("Build HTTP Request: ", err)
// 		return nil, err
// 	}

// 	// Add API key to the URL query
// 	for key, value := range config.APIKeys {
// 		req.URL.Query().Add(key, value)
// 	}

// 	for param, value := range apiReq.params {
// 		req.URL.Query().Add(param, value)
// 	}

// 	return req, nil
// }
