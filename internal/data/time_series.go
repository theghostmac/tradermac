package data

import (
	"database/sql"
	"encoding/json"
	"github.com/theghostmac/tradermac/internal/options"
	"log"
	"strconv"
	"time"
)

// ProcessTimeSeriesData processes the time series data and saves it to the database.
func ProcessTimeSeriesData(jsonData string, db *sql.DB) {
	var tsData TimeSeriesDailyData
	err := json.Unmarshal([]byte(jsonData), &tsData)
	if err != nil {
		log.Printf("Error unmarshalling time series data: %v", err)
		return
	}

	for dateStr, dailyPrice := range tsData.TimeSeriesDaily {
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			log.Printf("Error parsing date: %v", err)
			continue
		}

		price, err := strconv.ParseFloat(dailyPrice.Close, 64)
		if err != nil {
			log.Printf("Error parsing price: %v", err)
			continue
		}

		volume, err := strconv.ParseInt(dailyPrice.Volume, 10, 64)
		if err != nil {
			log.Printf("Error parsing volume: %v", err)
			continue
		}

		// Save the historical data.
		err = options.SaveHistoricalData(db, tsData.MetaData.Symbol, price, volume, date)
		if err != nil {
			log.Printf("Error saving historical data: %v", err)
			return
		}
	}
}
