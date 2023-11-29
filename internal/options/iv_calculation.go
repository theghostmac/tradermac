package options

import (
	"errors"
	"math"
)

const TOLERANCE = 1e-5
const MAXITERATIONS = 100

var sigma float64 = 0.2 // initial guess.

// CalculateIV calculates the implied volatility for a given option.
func CalculateIV(option Option, marketPrice, riskFreeRate, timeToExpiration float64) (float64, error) {
	for i := 0; i < MAXITERATIONS; i++ {
		price, err := optionPrice(option, sigma, riskFreeRate, timeToExpiration)
		if err != nil {
			return 0, err
		}

		vega := CalculateVega(option, option.UnderlyingPrice, riskFreeRate, sigma)
		if vega == 0 {
			return 0, errors.New("vega is zero")
		}

		diff := marketPrice - price
		if math.Abs(diff) < TOLERANCE {
			return sigma, nil
		}

		sigma += diff / vega
	}

	return 0, errors.New("failed to converge")
}

func optionPrice(option Option, sigma, riskFreeRate, timeToExpiration float64) (float64, error) {
	if option.Type == Call {
		return BlackScholesMertonCall(option.UnderlyingPrice, option.StrikePrice, sigma, riskFreeRate, timeToExpiration)
	} else if option.Type == Put {
		return BlackScholesMertonPut(option.UnderlyingPrice, option.StrikePrice, sigma, riskFreeRate, timeToExpiration)
	}

	return 0, errors.New("invalid option type")
}
