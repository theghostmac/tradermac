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
	UnderlyingPrice float64
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

type AssetType int

const (
	StockAsset AssetType = iota
	CommodityAsset
)

type CommodityType int

const (
	Gold CommodityType = iota
	Silver
)

// HistoricalData represents the historical data for a specific asset.
type HistoricalData struct {
	Date   time.Time
	Price  float64
	Volume int64
}

// ImpliedVolatility represents the implied volatility for a specific asset.
type ImpliedVolatility struct {
	Date         time.Time
	ImpliedVol   float64
	OptionSymbol string
}

// TradingSignal represents the trading signal for a specific asset.
type TradingSignal struct {
	Action       TradeAction // Buy or Sell
	SignalType   string      // "Overbought", or "Oversold"
	OptionSymbol string
}
