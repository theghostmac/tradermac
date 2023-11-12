package options

// NewPortfolio creates a new portfolio with an initial cash balance.
func NewPortfolio(initialCashBalance float64) *Portfolio {
	return &Portfolio{
		CashBalance:    initialCashBalance,
		OptionsHolding: make(map[OptionContract]uint),
	}
}
