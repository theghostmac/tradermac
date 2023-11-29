package trade

import (
	"database/sql"
	"github.com/theghostmac/tradermac/internal/options"
	"log"
	"time"
)

/*
The stdDevThreshold in a trading strategy is a critical parameter that determines
how sensitive your signal generation is to deviations in implied volatility (IV).
This threshold is typically set based on the specific characteristics of the market
you are trading in and the risk tolerance of your strategy. There isn't a
one-size-fits-all value; it requires some experimentation and backtesting to determine
what works best for your particular scenario. However, a common starting point in
many statistical models is to set this threshold to 2.
*/

const stdDevThreshold = 2

/*
A higher threshold (e.g., 2.5 or 3) will result in fewer, potentially more significant signals but might miss out on some trading opportunities.
A lower threshold (e.g., 1.5) will produce more signals, which could be beneficial in a highly volatile market, but may increase the risk of false positives.
*/

// generateSignals analyzes database IV data and underlying asset data to generate trading signals.
func generateSignals(db *sql.DB, optionSymbol string, startDate, endDate time.Time) ([]options.TradingSignal, error) {
	// Retrieve historical data for the underlying asset.
	historicalData, err := options.GetHistoricalData(db, optionSymbol, startDate, endDate)
	if err != nil {
		log.Printf("Error retrieving historical data: %v\n", err)
		return nil, err
	}

	// Retrieve implied volatility data.
	ivData, err := options.GetIVData(db, optionSymbol, startDate, endDate)
	if err != nil {
		log.Printf("Error retrieving IV data: %v\n", err)
		return nil, err
	}
	var signals []options.TradingSignal

	// Calculate the mean and standard deviation of the IV data.
	var ivValues []float64
	for _, iv := range ivData {
		ivValues = append(ivValues, iv.ImpliedVol)
	}

	meanIV, stdDevIV := calculateMeanAndStdDev(ivValues)

	// Determine if the asset is in an uptrend or downtrend.
	uptrend := isUptrend(historicalData, 50, 200)

	for _, iv := range ivData {
		if iv.ImpliedVol > meanIV+stdDevThreshold*stdDevIV {
			// Overbought condition.
			action := options.Sell
			if !uptrend {
				action = options.Buy
			}

			signals = append(signals, options.TradingSignal{
				OptionSymbol: iv.OptionSymbol,
				Action:       action,
				SignalType:   "Overbought",
			})
		} else if iv.ImpliedVol < meanIV-stdDevThreshold*stdDevIV {
			// Oversold condition.
			action := options.Buy
			if !uptrend {
				action = options.Sell
			}

			signals = append(signals, options.TradingSignal{
				OptionSymbol: iv.OptionSymbol,
				Action:       action,
				SignalType:   "Oversold",
			})
		}
	}

	return signals, nil
}

//// generateSignals analyzes database IV data and underlying asset data to generate trading signals.
//func generateSignals(ivData []options.ImpliedVolatility, historicalData []options.HistoricalData, stdDevThreshold float64) []options.TradingSignal {
//	var signals []options.TradingSignal
//
//	// Calculate the mean and standard deviation of the IV data.
//	var ivValues []float64
//	for _, iv := range ivData {
//		ivValues = append(ivValues, iv.ImpliedVol)
//	}
//
//	meanIV, stdDevIV := calculateMeanAndStdDev(ivValues)
//
//	// Determine if the asset is in an uptrend or downtrend.
//	uptrend := isUptrend(historicalData, 50, 200)
//
//	for _, iv := range ivData {
//		if iv.ImpliedVol > meanIV+stdDevThreshold*stdDevIV {
//			// Overbought condition.
//			action := options.Sell
//			if !uptrend {
//				action = options.Buy
//			}
//
//			signals = append(signals, options.TradingSignal{
//				OptionSymbol: iv.OptionSymbol,
//				Action:       action,
//				SignalType:   "Overbought",
//			})
//		} else if iv.ImpliedVol < meanIV-stdDevThreshold*stdDevIV {
//			// Oversold condition.
//			action := options.Buy
//			if !uptrend {
//				action = options.Sell
//			}
//
//			signals = append(signals, options.TradingSignal{
//				OptionSymbol: iv.OptionSymbol,
//				Action:       action,
//				SignalType:   "Oversold",
//			})
//		}
//	}
//
//	return signals
//}
