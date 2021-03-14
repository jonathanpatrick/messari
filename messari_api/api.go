package api

import (
	"encoding/json"
	"net/http"
	"time"
)

const messariBaseUrl = "https://data.messari.io/api/"
const assetPath = "v1/assets/"
const aggregatePath = "v2/assets?fields=id,name,slug,symbol,metrics/market_data/price_usd,metrics/market_data/volume_last_24_hours,metrics/market_data/percent_change_usd_last_24_hours,metrics/marketcap/current_marketcap_usd,metrics/misc_data/sectors,metrics/misc_data/tags&with-metrics&limit=500"

func GetAsset(asset string, target interface{}) error {
	client := getClient()
	url := messariBaseUrl + assetPath + asset + "/metrics"
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func GetAggregate(target interface{}) error {
	client := getClient()
	url := messariBaseUrl + aggregatePath
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func getClient() http.Client {
	return http.Client{
		Timeout: 5 * time.Second,
	}
}
