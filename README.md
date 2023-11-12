# TraderMac
TraderMac handles **fetching Options data** from a financial market, creating a **market** and **trading engine** 
for **options trading** between a buyer and a seller, use the Black-Scholes-Merton model to 
estimate price and buy the underlying asset to eliminate risks, [more features from research].

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
- [ ] Remodel the UnderlyingAsset
- [ ] Setup trader modeling and management | CRUD operations for trader.
- [ ] Write the logic to execute actual trades from a live source.
- [ ] Setup pipeline to simulate the execution of paper trades with a Brokerage API
- [ ] Eventually set pipeline to execute real trades with the Brokerage API
- [ ] Efficiently handle the TradeOrder OrderStatus property.
- [ ] Setup pipeline to external market data sources to get real-time/historical data for each underlying asset supported. (useful for pricing options, trading strategy, risk mgt).
- [ ] Add logging and maybe monitoring
- [ ] Write tests for everything until Coverage > 70%