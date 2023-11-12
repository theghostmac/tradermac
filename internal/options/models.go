package options

import "time"

// OptionType is the type of option, either "Call" or "Put".
type OptionType string

const (
	Call OptionType = "Call"
	Put  OptionType = "Put"
)

// Option represents an options contract.
type Option struct {
	Type            OptionType
	StrikePrice     float64
	ExpirationDate  time.Time
	UnderlyingAsset string // Symbol of identifier of underlying security or commodity.
}

// OptionContract represents a specific options contract.
type OptionContract struct {
	Option
	ContractID   uint
	ContractSize uint // optional.
}

// TradeAction represents the action taken on an options contract.
type TradeAction string

const (
	Buy  TradeAction = "Buy"
	Sell TradeAction = "Sell"
)

type TradeOrder struct {
	Action                TradeAction
	Option                Option
	Quantity              uint
	OrderPrice            float64
	OrderStatus           string // either "Open", "Filled", or "Cancelled"
	OrderTimeStamp        time.Time
	OrderExecutionDetails string
}
