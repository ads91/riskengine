# Risk Engine

## What is it?

An online back-end application for pricing financial instruments. The instruments are respresented as either a CSV or JSON file and return 


## Future enhancements

1. The market data environment is currently picked up from a static location (inside the repo). This should really be sent as part of the POST request to the application. In the simplest case, the environment is small - containing only one curve. However, depending on the use case i.e. the numbers of instruments to price, the complexity of those instruments, the market data environment could grow substantially and hence this will need to be taken into consideration.

2. Provide support for other trade representations, i.e. CSV, XML.



