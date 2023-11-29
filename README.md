# TraderMac

TraderMac is a sophisticated trading platform focused on options trading. 
It integrates various financial data sources to provide real-time analysis, 
trading signal generation, and portfolio management, specifically tailored for options trading.

## Features (see TODO section for more coming features)

- Modeling options contracts including calls and puts.
- Functions to calculate options value, risk metrics (delta, gamma, theta, vega), and execute trades.
- Integration with external data sources for real-time market data.
- A backtesting environment to simulate trading strategies on historical data.
- Portfolio management capabilities to track and manage trading activities.

## Getting Started

### Prerequisites

- Go (version 1.16 - 1.21)
- PostgreSQL
- Access to financial data APIs (e.g., Alpha Vantage)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/tradermac.git
   ```

2. Navigate to the project directory:
   ```sh
   cd tradermac
   ```

3. Install dependencies (if any):
   ```sh
   go get ./...
   ```

4. Set up the environment variables:
    - `DATABASE_URL`: Your PostgreSQL database URL.
    - `ALPHAVANTAGE_API_KEY`: Your API key for Alpha Vantage.

5. Run the application:
   ```sh
   go run main.go
   ```

### Configuration

- Configure database and API credentials in `.env` file or as environment variables.

## Usage

For now, I am using a dummy portfolio to backtest the model.

## TODO

- [x] Model an Options Contract
- [x] Implement data fetching from Alpha Vantage
- [x] Create database schema for options data
- [x] Develop functions for options valuation and risk management
- [x] Build a backtesting environment for strategy testing
- [x] Implement portfolio management features
- [ ] Integrate real-time data fetching from Yahoo Finance
- [ ] Refine `GetOptionBySymbol` function to ensure accurate data retrieval
- [ ] Implement comprehensive error handling and logging
- [ ] Optimize performance for high-frequency trading scenarios
- [ ] Conduct thorough testing of all components
- [ ] Deploy the application in a cloud environment
- [ ] Test in a live environment

## Resources

Resources I must remember to consume later live here.

1. [Awesome Systematic Trading](https://wangzhe3224.github.io/awesome-systematic-trading/#general-purpose)
2. 