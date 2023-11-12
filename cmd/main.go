package main

import (
	"fmt"
	"github.com/theghostmac/tradermac/internal/options"
	"time"
)

func main() {
	// Create a new portfolio with an initial cash balance.
	initialCashBalance := 10000.0
	portfolio := options.NewPortfolio(initialCashBalance)

	// Display initial portfolio information.
	fmt.Println("Initial Cash Balance: $", portfolio.CashBalance)
	fmt.Println("Initial Options Holding: No options initially")

	// Example trade order.
	expirationDate := time.Now()

	order := options.TradeOrder{
		Action: options.Buy,
		Option: options.Option{
			Type:            options.Call,
			StrikePrice:     150,
			ExpirationDate:  expirationDate,
			UnderlyingAsset: "AAPL",
		},
		Quantity:    5,
		OrderPrice:  10.0,
		OrderStatus: "Open",
		PortfolioID: 1,
	}

	// Execute the trade.
	err := options.ExecuteTrade(portfolio, order)
	if err != nil {
		fmt.Println("Trade execution failed: ", err)
		return
	}

	// Display updated portfolio information.
	fmt.Println("New Cash Balance: $", portfolio.CashBalance)
	fmt.Println("New Options Holding:")
	for contract, quantity := range portfolio.OptionsHolding {
		fmt.Printf("- %s: %d contracts\n", contract.Option, quantity)
	}
}
