package tests

import (
	"github.com/theghostmac/tradermac/internal/alphavantage"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestAPIClient_CallAlphaVantageAPI_Success(t *testing.T) {
	// Mocking a successful HTTP response from AlphaVantage.
	mockHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write([]byte(`{"symbol":"AAPL", "price":"123.45"}`))
	})

	mockServer := httptest.NewServer(mockHandler)
	defer mockServer.Close()

	// Creating a mock configuration
	actualAPIKey := os.Getenv("ALPHAVANTAGE_API_KEY")
	mockConfig := alphavantage.NewConfiguration(mockServer.URL).AddKey("apikey", actualAPIKey).CustomTimeout(5 * time.Second)

	// Creating a mock client
	mockClient := alphavantage.NewAPIClient(mockConfig)

	// Creating a mock request
	mockAPIRequest := alphavantage.NewAPIRequest().SetAPIReqPath("GET").SetAPIReqPath("/mockendpoint").SetAPIReqParam("symbol", "AAPL")

	// Creating a mock response
	mockAPIResponse, err := mockClient.CallAlphaVantageAPI(mockAPIRequest)

	// Asserting the result
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedResponse := `{"symbol":"AAPL", "price":"123.45"}`
	if mockAPIResponse != expectedResponse {
		t.Fatalf("Expected %s, got %s", expectedResponse, mockAPIResponse)
	}
}
