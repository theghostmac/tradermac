package options

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
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

// GetOptionBySymbol scrapes options data for a given symbol from Yahoo! Finance and returns the Option object for a given symbol.
func GetOptionBySymbol(symbol string) (Option, error) {
	url := "https://finance.yahoo.com/quote/" + symbol + "/options?p=" + symbol

	resp, err := http.Get(url)
	if err != nil {
		return Option{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Option{}, err
	}

	// Use regular express or HTML parsing library to extract data.
	match := regexp.MustCompile(``).FindSubmatch(body)
	if match == nil {

	}

	option := Option{
		// Values
	}

	return option, nil
}
