package options

import "errors"

var ErrInvalidOptionType = errors.New("invalid Option type")

var ErrInvalidOrderQuantity = errors.New("invalid Quantity in trade Order")

var ErrInvalidTradeAction = errors.New("invalid Trade Action")

var ErrInsufficientContract = errors.New("insufficient contracts to sell from the seller")

var ErrInsufficientFunds = errors.New("insufficient funds in the trader's portfolio")
