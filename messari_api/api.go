package api

import (
	"encoding/json"
	"net/http"
)

const messariBaseUrl = "https://data.messari.io/api/"
const assetPath = "v1/assets/"
const aggregatePath = "v2/assets"

func GetAsset(asset string, target interface{}) error {
	url := messariBaseUrl + assetPath + asset + "/metrics"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}
