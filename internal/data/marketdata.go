package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/theghostmac/tradermac/internal/alphavantage"
)

type MarketClient interface {
	GetTimeSeriesDaily(fn string, sb string, ak string) (string, error)
	GetCurrentStockPrice(fn string, sb string, ak string) (float64, error)
}

var _ MarketClient = &MarketData{}

// MarketData is responsible for retrieving market data.
type MarketData struct {
	alphavantageClient alphavantage.Client
}

// NewMarketData creates a new instance of MarketData.
func NewMarketData(alphavantageClient alphavantage.Client) *MarketData {
	return &MarketData{alphavantageClient: alphavantageClient}
}

// GetCurrentStockPrice retrieves the current stock price for a given symbol.
func (md *MarketData) GetCurrentStockPrice(fn string, sb string, ak string) (float64, error) {
	// Model the returned data against the response from AlphaVantage.
	var (
		path = fmt.Sprintf("function=%s&symbol=%s&apikey=%s", fn, sb, ak)
		apiReq *http.Request
		apiRes *http.Response
		err error
		headers = map[string]string{
			"Content-Type": "application/json",
		}
	)

	apiReq, err = md.alphavantageClient.MakeRequest(path, http.MethodGet, "", headers)
	if err != nil {
		log.Printf("Cannot make request: '%s'\n", err)
		return 0.0, err
	}

	apiRes, err = md.alphavantageClient.DoRequest(apiReq)
	if err != nil {
		log.Printf("Cannot do request: '%s'\n", err)
		return 0.0, err
	}

	defer apiRes.Body.Close()

	body, err := io.ReadAll(apiRes.Body)
	if err != nil {
		log.Printf("Error parsing to bytes: '%s'\n", string(body))
		return 0.0, err
	}

	// Unmarshall to a struct.
	fmt.Printf("Raw response: '\n%s\n'", string(body))

	// Fetch the byte data from io.ReadAll.
	var quote struct {
		Price float64 `json:"price"`
	}

	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Printf("Failed to parse the response: %v", err)
		return 0.0, err
	}

	return quote.Price, nil
}

func (md *MarketData) GetTimeSeriesDaily(fn string, sb string, ak string) (string, error) {
	var (
		path    = fmt.Sprintf("function=%s&symbol=%s&apikey=%s", fn, sb, ak)
		req     *http.Request
		res     *http.Response
		err     error
		headers = map[string]string{
			"Content-Type": "application/json",
		}
	)
	req, err = md.alphavantageClient.MakeRequest(path, http.MethodGet, "", headers)
	if err != nil {
		log.Printf("cannot make request:'%s\n'", err)
		return "", err
	}
	res, err = md.alphavantageClient.DoRequest(req)
	if err != nil {
		log.Printf("cannot do request:'%s\n'", err)
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error parsing to bytes: '%s\n'", string(body))
		return "", err
	}
	/* you can unmarshal to a struct */
	fmt.Printf("raw response: '\n%s\n'", string(body))
	/* marshal later. am trying to get the raw data*/
	return string(body), nil
}