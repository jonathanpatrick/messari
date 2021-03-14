
## Assumptions

Limiting aggregate call results to assets with a marketcap > 20M


## Improvements
Below are some areas for improvement for this service if it were a more full-scale project, time, etc

###### API pagination on Aggregate assets call 

Currently due to the minimum market cap of 20M, pagination was not implemented. However an improvement to make this service more robust would be implementing a proper pagination system to properly handle a future scenario where more than 500 assets (exceeding the upper limit of the Assets API in a single call) and we need to perform multiple calls to get all of the data.

###### Dynamic field selection

Allowing the requester to specify the fields similar to the Messari API would be a nice quality of life improvement.

###### More robust config implementation

This implementation/testing did not hit any rate limiting or API key issues, but controlling and supplying keys or configurable business logic (like min Aggregate Marketcap) via config would be a better practice for maintainability.