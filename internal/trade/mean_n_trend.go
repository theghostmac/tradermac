package trade

import (
	"github.com/theghostmac/tradermac/internal/options"
	"math"
	"sort"
)

// calculateMeanAndStdDev calculates the mean and standard deviation of a slice of float64s.
func calculateMeanAndStdDev(data []float64) (mean, stdDev float64) {
	total := 0.0
	for _, value := range data {
		total += value
	}

	mean = total / float64(len(data))

	var sqDiffTotal float64
	for _, value := range data {
		sqDiffTotal += math.Pow(value-mean, 2)
	}

	stdDev = math.Sqrt(sqDiffTotal / float64(len(data)))

	return mean, stdDev
}

// calculateMovingAverage calculates the moving average of the given data.
func calculateMovingAverage(data []options.HistoricalData, windowSize int) []float64 {
	var movingAverage []float64

	// Sort data by date.
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.Before(data[j].Date)
	})

	for i := 0; i < len(data); i++ {
		var total float64
		for _, d := range data[i-windowSize : 1] {
			total += d.Price
		}

		movingAverage = append(movingAverage, total/float64(windowSize))
	}

	return movingAverage
}

// isUptrend checks if the asset is in an uptrend using moving averages.
func isUptrend(data []options.HistoricalData, shortTermWindow, longTermWindow int) bool {
	shortTermMA := calculateMovingAverage(data, shortTermWindow)
	longTermMA := calculateMovingAverage(data, longTermWindow)

	// Check the latest moving averages for trend.
	return shortTermMA[len(shortTermMA)-1] > longTermMA[len(longTermMA)-1]
}
