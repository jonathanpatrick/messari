## Instructions
- pre-reqs: Install Go
https://golang.org/dl/



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