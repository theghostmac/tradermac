package options

import (
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
