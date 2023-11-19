package options

import (
	"errors"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

// ------------> DEFINITIONS <-----------------

// CDF is the function that gives the probability that a random variable takes a value less than or
// equal to a given value. For a continuous random variable, it is typically represented by an integral.

/* BSM returns theoretical price of an option.
 * receives: currentStockPrice or underlyingAssetPrice
 * 			 strikePrice
 * 			 cumulativeDistributionFunction
 *			 timeToExpiration
 * 			 volatilityIndex of the underlying asset
 */

// BlackScholesMertonCall calculates the Black-Scholes price for a European Call Option.
func BlackScholesMertonCall(currentStockPrice, strikePrice, volatility, riskFreeRate, timeToExpiration float64) (theoreticalCallPrice float64, err error) {
	// Convert the timeToExpiration to years.
	T := timeToExpiration / 365

	// Calculate d1 and d2
	d1 := (math.Log(currentStockPrice/strikePrice) + (riskFreeRate+0.5*math.Pow(volatility, 2))*T) / (volatility * math.Sqrt(T))

	d2 := d1 - volatility*math.Sqrt(T)

	// Calculate call option price using Black-Scholes formula.
	theoreticalCallPrice = currentStockPrice*NormCumulativeDistributionFunction(d1) - strikePrice*math.Exp(-riskFreeRate*T)*NormCumulativeDistributionFunction(d2)

	return theoreticalCallPrice, nil
}

// BlackScholesMertonPut calculate the Black-Scholes price for a European Put Option.
func BlackScholesMertonPut(currentStockPrice, strikePrice, volatility, riskFreeRate, timeToExpiration float64) (theoreticalPutPrice float64, err error) {
	// Convert the timeToExpiration to years.
	T := timeToExpiration / 365

	// Calculate d1 and d2.
	d1 := (math.Log(currentStockPrice/strikePrice) + (riskFreeRate+0.5*math.Pow(volatility, 2))*T) / (volatility * math.Sqrt(T))

	d2 := d1 - volatility*math.Sqrt(T)

	// Calculate the put option price using Black-Scholes formula.
	theoreticalPutPrice = strikePrice*math.Exp(-riskFreeRate*T)*NormCumulativeDistributionFunction(-d2) - currentStockPrice*NormCumulativeDistributionFunction(-d1)

	return theoreticalPutPrice, nil
}

// NormCumulativeDistributionFunction calculates the cumulative standard normal probability distribution function.
func NormCumulativeDistributionFunction(x float64) float64 {
	// Create a standard normal distribution
	normalDist := distuv.Normal{
		Mu:    0, // mean
		Sigma: 1, // standard deviation
	}

	// Calculate and return the CDF at the specified value (x).
	return normalDist.CDF(x)
}

// CalculateVolatilityIndex calculates the historical volatility for the given price data.
func CalculateVolatilityIndex(priceData []float64) (volatility float64, err error) {
	if len(priceData) < 2 {
		return 0, errors.New("insufficient data for volatility calculation")
	}

	var logReturns []float64
	for i := 1; i < len(priceData); i++ {
		logReturn := math.Log(priceData[i] / priceData[i-1])
		logReturns = append(logReturns, logReturn)
	}

	mean, sum := 0.0, 0.0
	for _, lr := range logReturns {
		mean += lr
	}

	mean /= float64(len(logReturns))

	for _, lr := range logReturns {
		sum += math.Pow(lr-mean, 2)
	}

	variance := sum / float64(len(logReturns)-1)
	standardDeviation := math.Sqrt(variance)

	// Annualized the standard deviation to represent the annual volatility
	// Assuming 252 trading days in a year.
	volatility = standardDeviation * math.Sqrt(252)
	return volatility, nil
}
