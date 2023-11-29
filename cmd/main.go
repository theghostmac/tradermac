package main

import (
	"crypto/tls"
	"fmt"
	"github.com/theghostmac/tradermac/internal/backtest"
	"github.com/theghostmac/tradermac/internal/options"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/theghostmac/tradermac/internal/alphavantage"
	"github.com/theghostmac/tradermac/internal/data"
)

func main() {
	db := options.GetDBConnection()
	defer db.Close()

	// Create the ImpliedVolatility table.
	options.CreateImpliedVolatilityTable(db)

	// Create the HistoricalData table.
	options.CreateHistoricalDataTable(db)

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

	stockQuote, err := mkd.GetTimeSeriesDaily("GLOBAL_QUOTE", "IBM", APIKEY)
	if err != nil {
		log.Printf("cannot fetch time series daily: '%s\n'", err)
		return
	}
	/* print to stdout */
	fmt.Printf("global stock quote data: '%s\n'", stockQuote)

	time_series, err = mkd.GetTimeSeriesDaily("TIME_SERIES_DAILY", "IBM", APIKEY)
	if err != nil {
		log.Printf("cannot fetch time series daily: %s\n", err)
		return
	}

	// TODO: Thoroughly test ProcessTimeSeriesData with various data inputs to ensure it handles different scenarios correctly, including malformed data and edge cases.
	data.ProcessTimeSeriesData(time_series, db)

	// Define backtest parameters.
	startDate := time.Date(2023, 11, 20, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2023, 11, 28, 0, 0, 0, 0, time.UTC)
	symbol := "IBM"
	initialCapital := 10_000.0

	// Create a new backtest instance.
	bt := backtest.NewBacktest(db, startDate, endDate, symbol, initialCapital)

	// Run the backtest.
	bt.Run()
}
