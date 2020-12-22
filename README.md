# Risk Engine

## What is it?

An online pricing engine for financial instruments. Instruments are respresented as a JSON file and return a JSON with the calculated price.

## Usage

...

## Supported instruments

Below outlines the different financial instruments that are supported by the risk engine. Multiple instruments, different types, multiple of the same type or a combination of these is supported. The only requirement is that the top-level key for each instrument configuration is unique within each request. 

For each instrument type, it has an associated args key in its configuration, this specifies the model parameters required to price the given type of instrument.

The following JSON is a an example of a valid pricing request, consisting of two instrument configurations for instrument type "instrument_a" and one instrument configuration for instrument type "instrument_b", both of which are hypothetical instrument types.

```json
{
    ...,
    "instrument_a_001": {
        "type": "instrument_a",
        "args": {
            ...
        }
    },
    "instrument_a_002": {
        "type": "instrument_a",
        "args": {
            ...
        }
    },
    "instrument_b_001": {
        "type": "instrument_b",
        "args": {
            ...
        }
    },
    ...
}
```

The type key in an instrument's configuration acts as a switch within the risk engine to call the correct pricing analytics for that instrument. The following sections outline the specifics of the supported instruments; the type parameters of those instruments and the required args by the risk engine to price those instruments.

*Note: if an unrecognised instrument type exists in the pricing request, the return result will flag the instrument with an error. If the instrument pricing request has other valid requests, these will still price in accordance with the associated analytics.*

### European call option
**type: europeancall**

A vanilla european call option, which can apply to any product that has a zero-dimensional underlying price i.e. an FX rate, an equity price, a bond price etc. Its representation is as follows

```json
{
    ...,
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
    },
    ...
}
```

where

- **startprice** is the initial price of the underlying,
- **strike** is the agreed strik price of the option,
- **years** is the time until the option expires,
- **rate** is the rate in which we discount the expected pay-off,
- **vol** is the volatility of the underlying (to be calculated by the caller of the risk engine) and
- **paths** is the number of monte-carlo simulations of the option pay-off.

#### Model limitations

- Doesn't support underlyings of greater than zero dimensions i.e. a yield curve.
- Discounting is done using a flat rate, rather than a relevant yield curve.
- Doesn't derive volatility from historic prices, so is a user input.

### Bond
**type: bond**

A vanilla bond. Pricing of the bond is performed using a discount curve - which implies the outstanding payment schedule of the bond. Its representation is as follows

```json
{   
    ...,
    "bond_001": {
        "type": "bond",
        "args": {
            "coupon": "123.0",
            "curve" : "libor"
        }
    },
    ...
}
```

where

- **coupon** is the amount payable to the bond holder on each date specified in the bond's curve and
- **curve** is a valid curve name stored in the market data environment, this specifies the rates to convert to discounting factors to levy on the coupons and the dates associated with those rates.

#### Model limitations

- Doesn't include the face value of the bond in the pricing, only the outstanding coupon payments, as implied by the curve.

## Market data

For some instrument types, data required for pricing the instrument is specified in the market data environment ("env"). The market data environment is a JSON that's saved within the risk engine folder structure (./riskengine/data/env.json).

Top-level keys in the JSON refer to broad market data types i.e. curves, surfaces, etc. The pricing analytics layer accesses the market data by assuming a naming convention within the env. Therefore, if the exisiting naming structure is modified, it can potentially break existing analytics if not tested accordingly. 

Adding new levels to the env is done at the discretion of the developer. However, adhering to the existing structure is best practice i.e. if a new pricing curve is required, it should be assigned a name and saved under the existing curves top-level key.

Below we outline the supported market data types and their required structure. This evolves through time and is a function of what the pricing analytics layer deems necessary - so the structure of the env is a function of the pricing analytics and the pricing analytics is a function of the env.

*Note: if market data is calculated dynamically or needs to be refreshed, the newly generated env.json will need to be saved over the existing one and the risk engine application restarted in order to load the market data in to memory.*

### Curves
**key: curves**

A curve is a one-dimensional representation of dates (tenors) against rates. The tenors can be strings or numeric the rates are to be numeric. An example curve is as follows

```json
{   
    ...,
    "curves": {
        ...,
        "curve_a": {
            "tenors": ["1y", "2y", "3y", "4y", "5y"],
            "rates" : [0.15, 0.30, 0.35, 0.50, 0.75]
        },
        ...
    },
    ...
}
```

where 

- *curve_a* (can be any name) is a user-defined name that prescribes some meaning to the curve,
- *tenors* contains a list of the dates of the curve and
- *rates* is the list of rates which are sequantially associated with the tenors outlined above.

*Note: the length of the tenors list and rates list in a curve entry must be the same.*

## Future enhancements

1. The market data environment is currently picked up from a static location (inside the repo). This should really be parameterised as part of the POST request to the application. In the simplest case, the environment is small - containing only one curve. However, depending on the use case i.e. the number of instruments to price, the complexity of those instruments, the market data environment could grow substantially leading to latency issues.

2. Provide support for other trade representations, i.e. CSV, XML.