package options

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// NewPortfolio creates a new portfolio with an initial cash balance.
func NewPortfolio(initialCashBalance float64) *Portfolio {
	return &Portfolio{
		CashBalance:    initialCashBalance,
		OptionsHolding: make(map[OptionContract]uint),
	}
}

// PrintStatus prints the current status of the portfolio.
func (p *Portfolio) PrintStatus() {
	fmt.Printf("Current Portfolio Status:\n")
	fmt.Printf("Cash Balance: $%.2f\n", p.CashBalance)
	fmt.Println("Options Holdings:")
	for contract, quantity := range p.OptionsHolding {
		fmt.Printf("%v: %d\n", contract, quantity)
	}
}

// GetOptionBySymbol returns the Option object for a given symbol.
func GetOptionBySymbol(symbol string) (Option, error) {
	// TODO: implement this function to fetch Options from a data broker online.

	// Fetch options details from and external data source.
	// make HTTP requests and parse the response.

	resp, err := http.Get("")
	if err != nil {
		return Option{}, err
	}

	defer resp.Body.Close()

	var option Option
	if err := json.NewDecoder(resp.Body).Decode(&option); err != nil {
		return Option{}, err
	}

	return Option{}, nil
}
