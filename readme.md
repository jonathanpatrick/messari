

#Assumptions



#Improvements

API pagination on Aggregate assets call - Currently due to the minimum market cap of 20M, pagination was not implemented. However an improvement to make this service more robust would be implementing a proper pagination system to properly handle a future scenario where more than 500 assets (exceeding the upper limit of the Assets API in a single call) and we need to perform multiple calls to get all of the data.