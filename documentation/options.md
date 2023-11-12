# Definitions

## Strike Price
The strike price of an option is a fixed price at which the owner of the option can buy ("call" option)
or sell ("put" option) the underlying security of commodity.

It may be gotten by either of two ways:
- get the spot price (the current market price).
- fixed as a discount or a premium.

## Moneyness
Moneyness is the value of a financial contract if the contract settlement is financial.
It is the difference between the strike price of the option and the current trading price 
of the underlying security.

Moneyness is described using _in-the-money_, _at-the-money_, and _out-of-the-money_.
- a call option is _in-the-money_ if the strike price is below the market price of the underlying.
- a put option is _in-the-money_ if the strike price is above the market price of the underlying.
- a call or put option is _at-the-money_ if the underlying price and stock price are the equal or similar.
- a call option is _out-of-the-money_ if the strike price is above the market price.
- a put option is _out-of-the-money_ if the strike price is below the market price.

## Black-Scholes Model
The Black-Scholes Model provides a formula for calculating the theoretical price of a financial option, 
taking into account factors such as the current stock price, 
the option's strike price, the time until expiration, the risk-free interest rate, and the 
volatility of the underlying asset's returns.

The key components of the Black-Scholes Model include:
- Option Price (C or P): The theoretical price of a call (C) or put (P) option.
- Current Stock Price (S): The market price of the underlying asset.
- Strike Price (K): The fixed price at which the option holder can buy (in the case of a call option) or sell (in the case of a put option) the underlying asset. 
- Time to Expiration (T): The remaining time until the option contract expires. 
- Risk-Free Interest Rate (r): The interest rate with no risk of financial loss. 
- Volatility (σ): A measure of the variability of the underlying asset's returns over time.

