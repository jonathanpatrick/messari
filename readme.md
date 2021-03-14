## Instructions
- pre-reqs: Install Go
https://golang.org/dl/

- Building (from project root)
go build -v ./...

- Run (from project root)
go run cmd/server/main.go


## Sample Responses
GET
http://localhost/api/asset/btc
{
    "symbol": "BTC",
    "name": "Bitcoin",
    "slug": "bitcoin",
    "price": 60176.95532767127,
    "volume": 46748041423.250854,
    "24hr_change": -1.5128116869470678,
    "marketcap": 1118513475697.2324
}

GET
http://localhost/api/asset/aave
{
    "symbol": "AAVE",
    "name": "Aave",
    "slug": "aave",
    "price": 390.0963826992281,
    "volume": 415625040.3329533,
    "24hr_change": -5.976104694365728,
    "marketcap": 4833784440.392679
}

GET
http://localhost/api/aggregate?tags=DEFI
{
    "tags": "DEFI",
    "count": 11,
    "volume": 1059072192.9993216,
    "marketcap": 12303450152.417534,
    "24hr_change": -4.486490264220324
}

GET
http://localhost/api/aggregate?sector=Smart%20Contract%20Platforms
{
    "sector": "Smart Contract Platforms",
    "count": 27,
    "volume": 33737063802.698685,
    "marketcap": 277614308795.3996,
    "24hr_change": -2.1022172253569638
}


## Assumptions

- Limiting aggregate call results to assets with a marketcap > 20M
- Aggregate 24 hour change calculation was made through calculating yesterdays market cap (per asset), then comparing yesterday's cumulative aggregate vs. today's cumulative aggregate.

## Improvements
Below are some areas for improvement for this service if it were a more full-scale project, time, etc

###### Testing
This project includes a couple tests for a helper function used. Adding better unit test coverage and integration tests would help maintain the integrity of the service.

###### API pagination on Aggregate assets call 

Currently due to the minimum market cap of 20M, pagination was not implemented. However an improvement to make this service more robust would be implementing a proper pagination system to properly handle a future scenario where more than 500 assets (exceeding the upper limit of the Assets API in a single call) and we need to perform multiple calls to get all of the data.

###### Dynamic field selection

Allowing the requester to specify the fields similar to the Messari API would be a nice quality of life improvement.

###### Dynamic config implementation for server vars

This implementation/testing did not hit any rate limiting or API key issues, but controlling and supplying keys, server variables, or configurable business logic (like min Aggregate Marketcap) via config would be a best practice for maintainability.