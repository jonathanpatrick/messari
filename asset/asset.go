package asset

func MapAssetResponse(data AssetData) AssetResponse {
	return AssetResponse{
		Symbol:               data.Asset.Symbol,
		Name:                 data.Asset.Name,
		Slug:                 data.Asset.Slug,
		Price:                data.Asset.MarketData.Price,
		Volume:               data.Asset.MarketData.Volume,
		TwentyFourHourChange: data.Asset.MarketData.TwentyFourHourChange,
		Marketcap:            data.Asset.Marketcap.Marketcap,
	}
}

// Struct for output response
type AssetResponse struct {
	Symbol               string  `json:"symbol"`
	Name                 string  `json:"name"`
	Slug                 string  `json:"slug"`
	Price                float64 `json:"price"`
	Volume               float64 `json:"volume"`
	TwentyFourHourChange float64 `json:"24hr_change"`
	Marketcap            float64 `json:"marketcap"`
}

// Structs for capturing messari response
type AssetData struct {
	Asset struct {
		Symbol     string `json:"symbol,omitempty"`
		Name       string `json:"name,omitempty"`
		Slug       string `json:"slug,omitempty"`
		MarketData struct {
			Price                float64 `json:"price_usd,omitempty"`
			Volume               float64 `json:"volume_last_24_hours,omitempty"`
			TwentyFourHourChange float64 `json:"percent_change_usd_last_24_hours,omitempty"`
		} `json:"market_data"`
		Marketcap struct {
			Marketcap float64 `json:"current_marketcap_usd,omitempty"`
		} `json:"marketcap"`
	} `json:"data"`
}
