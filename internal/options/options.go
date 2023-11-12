package options

// CalculateOptionValue calculates the value of an options position based on market conditions.
func CalculateOptionValue(option Option, currentPrice float64) (float64, error) {
	if option.Type == Call {
		return calculateCallOptionValue(option, currentPrice), nil
	} else if option.Type == Put {
		return calculatePutOptionValue(option, currentPrice), nil
	}
	return 0, ErrInvalidOptionType
}

func calculatePutOptionValue(option Option, currentPrice float64) float64 {
	// the value of a put option is the option's strike price minus the current price.
	intrinsicValue := option.StrikePrice - currentPrice

	if intrinsicValue < 0 {
		return 0
	}

	return intrinsicValue
}

func calculateCallOptionValue(option Option, currentPrice float64) float64 {
	// the value of a call option is the current price minus the option's strike price.
	intrinsicValue := currentPrice - option.StrikePrice

	if intrinsicValue < 0 {
		return 0
	}

	return intrinsicValue
}

// ValidateTradeOrders checks if a trade order is valid based on the given option and market conditions.
func ValidateTradeOrders(order TradeOrder, currentPrice float64) error {
	if order.Quantity <= 0 {
		return ErrInvalidOrderQuantity
	}

	if order.Action == Buy {
		// TODO:
		// - check if buyer has sufficient funds.
		// - integrate the trader's portfolio.
		// - execute buy.
		return nil
	} else if order.Action == Sell {
		// TODO:
		// - check if the seller has enough options contract to sell.
		// - integrate the trader's portfolio.
		// - execute sell.
		return nil
	}

	return ErrInvalidTradeAction
}

//// ExecuteTrade executes a trade order and returns the resulting position.
//func ExecuteTrade(currentPosition uint, order TradeOrder) (uint, error) {
//	if order.Action == Buy {
//		return currentPosition + order.Quantity, nil
//	} else if order.Action == Sell {
//		// Check if seller has enough contracts to sell.
//		if currentPosition < order.Quantity {
//			return currentPosition, ErrInsufficientContract
//		}
//		return currentPosition - order.Quantity, nil
//	}
//
//	return currentPosition, ErrInvalidTradeAction
//}

// ExecuteTrade executes a trade order and updates the trader's portfolio.
func ExecuteTrade(portfolio *Portfolio, order TradeOrder) error {
	if order.Action == Buy {
		cost := order.OrderPrice * float64(order.Quantity)

		// Check if the trader has sufficient funds.
		if cost > portfolio.CashBalance {
			return ErrInsufficientFunds
		}

		// Update the cash balance and options holding.
		portfolio.CashBalance -= cost
		contract := OptionContract{
			Option:       order.Option,
			ContractID:   order.PortfolioID, // TODO: just for simplicity, change later.
			ContractSize: order.Quantity,
		}
		portfolio.OptionsHolding[contract] += order.Quantity
	} else if order.Action == Sell {
		// Check if the seller has enough contracts to sell.
		contract := OptionContract{
			Option:     order.Option,
			ContractID: order.PortfolioID, // TODO: just for simplicity, change later.
		}

		if order.Quantity > portfolio.OptionsHolding[contract] {
			return ErrInsufficientContract
		}

		// Update the cash balance and options holding.
		revenue := order.OrderPrice * float64(order.Quantity)
		portfolio.CashBalance += revenue
		portfolio.OptionsHolding[contract] -= order.Quantity
	}

	return nil
}
