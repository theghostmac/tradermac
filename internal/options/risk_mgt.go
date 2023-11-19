package options

import (
	"math"
	"time"
)

// CalculateDelta measures the rate of change of the option price against the changes in the underlying asset's price.
func CalculateDelta(option Option, currentStockPrice, riskFreeRate, volatilityIndex float64) float64 {
	T := option.ExpirationDate.Sub(time.Now()).Hours() / (24 * 365)
	d1 := (math.Log(currentStockPrice/option.StrikePrice) + (riskFreeRate+0.5*math.Pow(volatilityIndex, 2))*T) / (volatilityIndex * math.Sqrt(T))

	if option.Type == Call {
		return NormCumulativeDistributionFunction(d1)
	} else if option.Type == Put {
		return NormCumulativeDistributionFunction(-d1)
	}

	return 0
}

// CalculateGamma measures the rate of change in Delta with respect to underlying asset's price.
func CalculateGamma(option Option, currentStockPrice, riskFreeRate, volatilityIndex float64) float64 {
	T := option.ExpirationDate.Sub(time.Now()).Hours() / (24 * 365)
	d1 := (math.Log(currentStockPrice/option.StrikePrice) + (riskFreeRate+0.5*math.Pow(volatilityIndex, 2))*T) / (volatilityIndex * math.Sqrt(T))
	phi := math.Exp(-math.Pow(d1, 2)/2) / math.Sqrt(2*math.Pi)

	return phi / (currentStockPrice * volatilityIndex * math.Sqrt(T))
}

// CalculateTheta measures the rate of change of the option price with respect to the passage of time.
func CalculateTheta(option Option, currentStockPrice float64, riskFreeRate float64, volatility float64) float64 {
	T := option.ExpirationDate.Sub(time.Now()).Hours() / (24 * 365)
	d1 := (math.Log(currentStockPrice/option.StrikePrice) + (riskFreeRate+0.5*math.Pow(volatility, 2))*T) / (volatility * math.Sqrt(T))
	d2 := d1 - volatility*math.Sqrt(T)
	phi := math.Exp(-math.Pow(d1, 2)/2) / math.Sqrt(2*math.Pi)

	commonTerm := -(currentStockPrice * phi * volatility) / (2 * math.Sqrt(T))
	if option.Type == Call {
		return commonTerm - riskFreeRate*option.StrikePrice*math.Exp(-riskFreeRate*T)*NormCumulativeDistributionFunction(d2)
	} else if option.Type == Put {
		return commonTerm + riskFreeRate*option.StrikePrice*math.Exp(-riskFreeRate*T)*NormCumulativeDistributionFunction(-d2)
	}
	return 0
}

// CalculateVega measures the rate of change of the option price with respect to changes in the volatility of the underlying asset.
func CalculateVega(option Option, currentStockPrice float64, riskFreeRate float64, volatility float64) float64 {
	T := option.ExpirationDate.Sub(time.Now()).Hours() / (24 * 365)
	d1 := (math.Log(currentStockPrice/option.StrikePrice) + (riskFreeRate+0.5*math.Pow(volatility, 2))*T) / (volatility * math.Sqrt(T))
	phi := math.Exp(-math.Pow(d1, 2)/2) / math.Sqrt(2*math.Pi)

	return currentStockPrice * phi * math.Sqrt(T)
}
