# TraderMac

TraderMac handles **fetching Options data** from a financial market, creating a **market** and **trading engine** 
for **options trading** between a buyer and a seller, use the Black-Scholes-Merton model to 
estimate price and buy the underlying asset to eliminate risks,[more features from research]

## Run

For now, manually create a portfolio to test the trading validation and execution:

```shell
go run cmd/main.go
```

Result:

```shell
Initial Cash Balance: $ 10000
Initial Options Holding: No options initially
New Cash Balance: $ 9950
New Options Holding:
- {Call 150.00 2023-11-12 14:41:33.476395 +0100 WAT m=+0.000110918 AAPL}: 5 contracts
```

## TODO:

- [x] Model an Options Contract
- [x] Calculate Options value
- [x] Validate a Trade Action
- [x] Execute Trade
- [x] Document Learnings
- [x] Test with main function.
- [x] Select brokerage API to get data from. AlphaVantage selected.
- [x] Setup data fetching with API requests from AlphaVantage.co
- [ ] Parse data to JSON and utilize it.
- [ ] Use data in BlackScholes implementation and relevant places.
- [ ] Connect BlackScholes implementation to everything that needs vending.
- [ ] Remodel the UnderlyingAsset
- [ ] Setup trader modeling and management | CRUD operations for trader.
- [ ] Write the logic to execute actual trades from a live source.
- [ ] Setup pipeline to simulate the execution of paper trades with a Brokerage API
- [ ] Eventually set pipeline to execute real trades with the Brokerage API
- [ ] Efficiently handle the TradeOrder OrderStatus property.
- [x] Setup pipeline to external market data sources to get real-time/historical data for each underlying asset supported. (useful for pricing options, trading strategy, risk mgt).
- [ ] Add logging and maybe monitoring
- [ ] Write tests for everything until Coverage > 70%

# Time Series

```shell
'tsd:'
{
    "Meta Data": {
        "1. Information": "Daily Prices (open, high, low, close) and Volumes",
        "2. Symbol": "IBM",
        "3. Last Refreshed": "2023-11-21",
        "4. Output Size": "Compact",
        "5. Time Zone": "US/Eastern"
    },
    "Time Series (Daily)": {
        "2023-11-21": {
            "1. open": "154.6000",
            "2. high": "154.6600",
            "3. low": "153.5100",
            "4. close": "153.9100",
            "5. volume": "2859508"
        },
        "2023-11-20": {
            "1. open": "152.5100",
            "2. high": "154.6800",
            "3. low": "152.3500",
            "4. close": "154.3500",
            "5. volume": "3658936"
        }
         ...
```

# Quote Response

```shell
 request took '220.170083ms
' to process
raw response: '
{
    "Global Quote": {
        "01. symbol": "IBM",
        "02. open": "154.6000",
        "03. high": "154.6600",
        "04. low": "153.5100",
        "05. price": "153.9100",
        "06. volume": "2859508",
        "07. latest trading day": "2023-11-21",
        "08. previous close": "154.3500",
        "09. change": "-0.4400",
        "10. change percent": "-0.2851%"
    }
}
```
