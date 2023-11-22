package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/theghostmac/tradermac/internal/alphavantage"
	"github.com/theghostmac/tradermac/internal/data"
)

func main() {
	var (
		httpTransport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		httpClient = &http.Client{
			Transport: httpTransport,
			Timeout:   time.Second * 15,
		}
	)

	var APIKEY = os.Getenv("ALPHAVANTAGE_MARKET_API_KEY")

	avc := alphavantage.NewRequestClient("https://www.alphavantage.co/query?", httpClient)
	mkd := data.NewMarketData(avc)
	time_series, err := mkd.GetTimeSeriesDaily("TIME_SERIES_DAILY", "IBM", APIKEY)
	if err != nil {
		log.Printf("cannot fetch time series daily: '%s\n'", err)
		return
	}
	/* print to stdout */
	fmt.Printf("time series daily: '%s\n'", time_series)

	stock_quote, err := mkd.GetTimeSeriesDaily("GLOBAL_QUOTE", "IBM", APIKEY)
	if err != nil {
		log.Printf("cannot fetch time series daily: '%s\n'", err)
		return
	}
	/* print to stdout */
	fmt.Printf("global stock quote data: '%s\n'", stock_quote)
}
