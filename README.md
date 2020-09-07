# FundSearch

FundSearch will run as a microservice using http to comunicate.

## Build:
`go run ./`

## Usage:
`localhost:10000/api/v1/fund/{FundName}`

Will bring back all holdings for the given fund, e.g.

`curl --location --request GET 'localhost:10000/api/v1/fund/Fund D'`

or for the specific scenario given:

`curl --location --request GET 'localhost:10000/api/v1/fund/Ethical Global Fund'`

By default it will use the file given for this task, however you can also pass your own funds file:

`curl --location --request GET 'localhost:10000/api/v1/fund/Ethical Global Fund' \
--form 'file=@/{PROJECT_ROOT}/examples/sample-funds2.json'`

Note that in the json passed in:  
* "Fund C" has been replaced by "Ethical Fund C"
* The troublesome holding of "Golden Gadgets" has been replaced by "Biodegradable Gadgets"  

You can see this by running:  
`curl --location --request GET 'localhost:10000/api/v1/fund/Ethical Fund C' \
--form 'file=@/{PROJECT_ROOT}/examples/sample-funds2.json'`

## Test:
`go test -v ./...`