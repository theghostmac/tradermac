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
- [ ] 