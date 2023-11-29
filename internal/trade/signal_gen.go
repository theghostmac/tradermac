package trade

import "github.com/theghostmac/tradermac/internal/options"

// generateSignals analyzes IV data and underlying asset data to generate trading signals.
func generateSignals(ivData []options.ImpliedVolatility, historicalData []options.HistoricalData, stdDevThreshold float64) []options.TradingSignal {
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

	return signals
}
