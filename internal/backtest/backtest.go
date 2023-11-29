package backtest

import (
	"database/sql"
	"fmt"
	"github.com/theghostmac/tradermac/internal/options"
	"github.com/theghostmac/tradermac/internal/trade"
	"log"
	"time"
)

type Backtest struct {
	DB             *sql.DB
	StartDate      time.Time
	EndDate        time.Time
	OptionSymbol   string
	InitialCapital float64
	Portfolio      *options.Portfolio
}

// NewBacktest creates a new Backtest environment.
func NewBacktest(db *sql.DB, startDate, endDate time.Time, symbol string, initialCapital float64) *Backtest {
	return &Backtest{
		DB:             db,
		StartDate:      startDate,
		EndDate:        endDate,
		OptionSymbol:   symbol,
		InitialCapital: initialCapital,
		Portfolio:      options.NewPortfolio(initialCapital),
	}
}

// Run starts the backtesting process.
func (b *Backtest) Run() {
	// Iterate over each day in the date range.
	for date := b.StartDate; !date.After(b.EndDate); date = date.AddDate(0, 0, 1) {
		// Fetch historical data for the day.
		historicalData, err := options.GetHistoricalData(b.DB, b.OptionSymbol, date, date)
		if err != nil {
			log.Printf("Error fetching historical data: %v", err)
			continue
		}

		// Generate trading signals for the day.
		signals, err := trade.GenerateSignals(b.DB, b.OptionSymbol, date, date)
		if err != nil {
			log.Printf("Error generating signals: %v", err)
			continue
		}

		// Process each trading signal.
		for _, signal := range signals {
			// Execute hypothetical trades and update the portfolio.
			b.ProcessSignal(signal, historicalData)
		}

		// Record the portfolio status at the end of the day.
		b.RecordPortfolioStatus(date)
	}

	// After backtesting, print the final portfolio status.
	b.Portfolio.PrintStatus()
}

// ProcessSignal processes a trading signal and updates the portfolio.
func (b *Backtest) ProcessSignal(signal options.TradingSignal, historicalData []options.HistoricalData) {
	if len(historicalData) == 0 {
		fmt.Println("No historical data available for trading signal.")
		return
	}

	// Fetch the Option details based on the OptionSymbol in the signal.
	option, err := options.GetOptionBySymbol(signal.OptionSymbol)
	if err != nil {
		log.Printf("Error fetching option details for symbol %s: %v", signal.OptionSymbol, err)
		return
	}
	currentPrice := option.UnderlyingPrice

	// assume historicalPrice is current price.
	currentPrice = historicalData[len(historicalData)-1].Price

	// Define a hypothetical quantity for the trade.
	quantity := uint(10) // can be dynamic value based on trading strategy.

	// Define the trade logic based on the signal.
	switch signal.Action {
	case options.Buy:
		// Calculate the total cost of the trade.
		totalCost := currentPrice * float64(quantity)

		// Check if there is enough cash to buy.
		if totalCost > b.Portfolio.CashBalance {
			fmt.Println("Insufficient funds to execute buy order.")
			return
		}

		// Update the portfolio: deduct cash and add option holding.
		b.Portfolio.CashBalance -= totalCost
		optionContract := options.OptionContract{
			Option:       option,
			ContractSize: quantity,
		}
		b.Portfolio.OptionsHolding[optionContract] += quantity

	case options.Sell:
		// Define the option contract to sell.
		optionContract := options.OptionContract{
			Option:       option,
			ContractSize: quantity,
		}

		// Check if the portfolio has enough of the option to sell.
		if b.Portfolio.OptionsHolding[optionContract] < quantity {
			fmt.Println("Insufficient options to execute sell order.")
			return
		}

		// Calculate the total revenue from the sale.
		totalRevenue := currentPrice * float64(quantity)

		// Update the portfolio: add cash and reduce option holding.
		b.Portfolio.CashBalance += totalRevenue
		b.Portfolio.OptionsHolding[optionContract] -= quantity
	default:
		fmt.Println("Unrecognized trade action.")
	}
}

// RecordPortfolioStatus records the status of the portfolio.
func (b *Backtest) RecordPortfolioStatus(date time.Time) {
	// Record the portfolio status for the given date
	// This can be as simple as logging the portfolio status or as complex as saving it to a file or database
	fmt.Printf("Portfolio status on %s: %+v\n", date.Format("2006-01-02"), b.Portfolio)

}
