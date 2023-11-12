package options

import (
	"fmt"
	"time"
)

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

// String method for Option to specify how it should be printed.
func (o Option) String() string {
	return fmt.Sprintf("{%s %.2f %s %s}", o.Type, o.StrikePrice, o.ExpirationDate, o.UnderlyingAsset)
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
	PortfolioID           uint
	OrderExecutionDetails string
}

// Portfolio represents the trader's portfolio.
type Portfolio struct {
	CashBalance    float64
	OptionsHolding map[OptionContract]uint
}
