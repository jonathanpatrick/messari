package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const messariBaseUrl = "https://data.messari.io/api/"
const assetPath = "v1/assets/"
const aggregatePath = "v2/assets?fields=id,name,slug,symbol,metrics/market_data/price_usd,metrics/market_data/volume_last_24_hours,metrics/market_data/percent_change_usd_last_24_hours,metrics/marketcap/current_marketcap_usd,metrics/misc_data/sectors,metrics/misc_data/tags&with-metrics&limit=100"

func GetAsset(asset string, target interface{}) error {
	url := messariBaseUrl + assetPath + asset + "/metrics"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func GetAggregate(tags string, sector string, target interface{}) error {
	queryParam := ""
	// Assumption: we are only allowing for tags or sector.
	// In the case both are provided, apply tags query param
	if tags != "" {
		queryParam = "&tags=" + tags
	} else if sector != "" {
		queryParam = "&sector=" + sector
	}
	url := messariBaseUrl + aggregatePath + queryParam
	fmt.Println("Printing URL: " + url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	return json.NewDecoder(resp.Body).Decode(target)
}
