# Risk Engine

## What is it?

An online back-end application for pricing financial instruments. The instruments are respresented as a JSON file and return a JSON with the calculated price.

## Supported instruments

### European call option

This instrument type is a simple vanilla european call option, which can apply to any product that has a zero-dimensional underlying price i.e. an FX rate, an equity price, a bond price etc. Its representation is as follows

```json
{
    "ID001": {
            "type": "europeancall",
            "args": {
                "startprice": "100.0",
                "strike"    : "125.0",
                "years"     : "5",
                "rate"      : "0.025",
                "vol"       : "1.0",
                "paths"     : "100000"
            }
        }
}
```

where

- *startprice* is the initial price of the underlying,
- *strike* is the agreed strik price of the option,
- *years* is the time until the option expires,
- *rate* is the rate in which we discount the expected pay-off,
- *vol* is the volatility of the underlying (to be calculated by the caller of the risk engine) and
- *paths* is the number of monte-carlo simulations of the option pay-off.

#### Model limitations

- Doesn't support underlyings of greater than zero dimensions i.e. a yield curve.
- Discounting is done using a flat rate, rather than a relevant yield curve.
- Doesn't derive volatility from historic prices, so is a user input.

## Future enhancements

1. The market data environment is currently picked up from a static location (inside the repo). This should really be sent as part of the POST request to the application. In the simplest case, the environment is small - containing only one curve. However, depending on the use case i.e. the numbers of instruments to price, the complexity of those instruments, the market data environment could grow substantially and hence this will need to be taken into consideration.

2. Provide support for other trade representations, i.e. CSV, XML.
