# Risk Engine

## What is it?

An online pricing engine for financial instruments. Instruments are respresented as a JSON file and return a JSON with the calculated price.

## Usage

...

## Supported instruments

Below outlines the different financial instruments that are supported by the risk engine. Multiple instruments, different types, multiple of the same type or a combination of these is supported. The only requirement is that the top-level key for each instrument configuration is unique within each request. 

For each instrument type, it has an associated args key in its configuration, this specifies the model parameters required to price the given type of instrument.

The following JSON is a an example of a valid pricing request, consisting of two instrument configurations for instrument type "INSTRUMENT_A" and one instrument configuration for instrument type "INSTRUMENT_B", both of which are hypothetical instrument types.

```json
{
    "INSTRUMENT_A_001": {
            "type": "INSTRUMENT_A",
            "args": {
                ...
            }
        },
    "INSTRUMENT_A_002": {
            "type": "INSTRUMENT_A",
            "args": {
                ...
            }
        },
    "INSTRUMENT_B_001": {
            "type": "INSTRUMENT_B",
            "args": {
                ...
            }
        }
}
```

The type key in an instrument's configuration acts as a switch within the risk engine to call the correct pricing analytics for that instrument. The following sections outline the specifics of the supported instruments; the type parameters of those instruments and the required args by the risk engine to price those instruments.

*Note: if an unrecognised instrument type exists in the pricing request, the return result will flag the instrument with an error. If the instrument pricing request has other valid requests, these will still price in accordance with the associated analytics.*

### European call option
**type: europeancall**

A vanilla european call option, which can apply to any product that has a zero-dimensional underlying price i.e. an FX rate, an equity price, a bond price etc. Its representation is as follows

```json
{
    "europeancall_001": {
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

### Bond
**type: bond**

A vanilla bond. Pricing of the bond is performed using a discount curve - which implies the outstanding length of the bond.

```json
{    
    "bond_001": {
        "type": "bond",
        "args": {
            "coupon": "123.0",
            "curve" : "libor"
        }
    }
}
```

where

- *coupon* is the amount payable to the bond holder on each date specified in the bond's curve
- *curve* is a valid curve name stored in the market data environment, this specified the discounting curve to apply.

## Future enhancements

1. The market data environment is currently picked up from a static location (inside the repo). This should really be sent as part of the POST request to the application. In the simplest case, the environment is small - containing only one curve. However, depending on the use case i.e. the number of instruments to price, the complexity of those instruments, the market data environment could grow substantially leading to latency issues.

2. Provide support for other trade representations, i.e. CSV, XML.
