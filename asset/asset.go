package asset

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/jonathanpatrick/messari/messari_api"
)

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetData := &AssetData{}
	err := api.GetAsset(vars["asset"], assetData)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := json.Marshal(mapAssetResponse(*assetData))
	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(resp))
}

// Struct for output response
type AssetResponse struct {
	Symbol               string  `json:"symbol"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	Price                float64 `json:"price"`
	Volume               float64 `json:"volume"`
	TwentyFourHourChange float64 `json:"24hr change"`
	Marketcap            float64 `json:"marketcap"`
	EthChange            float64 `json:"percent_change_eth_last_24_hours"`
}

func mapAssetResponse(data AssetData) AssetResponse {
	return AssetResponse{
		Symbol:               data.Asset.Symbol,
		Name:                 data.Asset.Name,
		Slug:                 data.Asset.Slug,
		Price:                data.Asset.MarketData.Price,
		Volume:               data.Asset.MarketData.Volume,
		TwentyFourHourChange: data.Asset.MarketData.TwentyFourHourChange,
		Marketcap:            data.Asset.Marketcap.Marketcap,
		EthChange:            data.Asset.MarketData.EthChange,
	}
}

// Structs for capturing messari response
type AssetData struct {
	Asset Asset `json:"data"`
}

type Asset struct {
	Symbol     string          `json:"symbol"`
	Name       string          `json:"name"`
	Slug       string          `json:"slug"`
	MarketData AssetMarketData `json:"market_data"`
	Marketcap  AssetMarketcap  `json:"marketcap"`
}

type AssetMarketData struct {
	Price                float64 `json:"price_usd"`
	Volume               float64 `json:"volume_last_24_hours"`
	TwentyFourHourChange float64 `json:"percent_change_usd_last_24_hours"`
	EthChange            float64 `json:"percent_change_eth_last_24_hours"`
}

type AssetMarketcap struct {
	Marketcap float64 `json:"current_marketcap_usd"`
}
