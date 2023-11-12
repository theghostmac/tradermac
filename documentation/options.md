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
The Black-Scholes Model is a mathematical model commonly used for pricing European-style options. Developed by economists Fischer Black and Myron Scholes in 1973, along with the assistance of Robert Merton, the model has been influential in the field of financial economics.

The Black-Scholes Model provides a formula for calculating the theoretical price of a financial option, taking into account factors such as the current stock price, the option's strike price, the time until expiration, the risk-free interest rate, and the volatility of the underlying asset's returns.

The key components of the Black-Scholes Model include:

1. **Option Price (C or P):** The theoretical price of a call (C) or put (P) option.

2. **Current Stock Price (S):** The market price of the underlying asset.

3. **Strike Price (K):** The fixed price at which the option holder can buy (in the case of a call option) or sell (in the case of a put option) the underlying asset.

4. **Time to Expiration (T):** The remaining time until the option contract expires.

5. **Risk-Free Interest Rate (r):** The interest rate with no risk of financial loss.

6. **Volatility (σ):** A measure of the variability of the underlying asset's returns over time.

The Black-Scholes formula for a call option (C) is:

\[ C = S_0 \cdot N(d_1) - K \cdot e^{-rT} \cdot N(d_2) \]

And for a put option (P):

\[ P = K \cdot e^{-rT} \cdot N(-d_2) - S_0 \cdot N(-d_1) \]

Where:
- \( N(x) \) is the cumulative distribution function of the standard normal distribution.
- \( d_1 = \frac{\ln\left(\frac{S_0}{K}\right) + \left(r + \frac{\sigma^2}{2}\right)T}{\sigma\sqrt{T}} \)
- \( d_2 = d_1 - \sigma\sqrt{T} \)

The Black-Scholes Model assumes that financial markets are efficient, and it has been widely used for options pricing, despite its simplifying assumptions and limitations, such as the assumption of constant volatility and no transaction costs.