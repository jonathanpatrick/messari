package aggregate

import (
	"errors"
	"strings"
)

func ProcessAggregateResponse(tags string, sector string, aggregateData AggregateData) (AggregateDataResponse, error) {
	if tags != "" {
		// Assumption: We are only allowing param for filtering
		// In the case both sector and tags are populated, we are defaulting
		// to using tags and setting sector to empty
		sector = ""
	}
	var resp AggregateDataResponse
	resp.Tags = tags
	resp.Sector = sector
	var cumulativeMarketCapYesterday float64
	for _, item := range aggregateData.Data {
		// Skip marketcaps less than 20M
		if item.Metrics.Marketcap.Marketcap < 20000000 {
			// since the messari API defaults to a desc marketcap sort
			// once we hit our first case < 20M, we can exit this loop
			break
		}
		if tags != "" && !contains(item.Metrics.MiscData.Tags, tags) {
			// skip asset missing tag
			continue
		} else if sector != "" && !contains(item.Metrics.MiscData.Sectors, sector) {
			// skip asset missing sector
			continue
		}
		resp.Count++
		resp.Volume += item.Metrics.MarketData.Volume
		resp.Marketcap += item.Metrics.Marketcap.Marketcap

		cumulativeMarketCapYesterday += adjustValueForYesterday(item.Metrics.Marketcap.Marketcap,
			item.Metrics.MarketData.TwentyFourHourChange)
	}
	if resp.Count == 0 {
		return resp, errors.New("No data retrieved, please validate your query parameters")
	}
	// Determine cumulative 24 hr change by comparing yesterdays cumulative
	// market cap to current market cap
	// Assumption: we want our aggregate 24 hour change to be properly weighted by market cap
	resp.TwentyFourHourChange = (resp.Marketcap - cumulativeMarketCapYesterday) / cumulativeMarketCapYesterday * 100 // percent formatting

	return resp, nil
}

// Struct for output response
type AggregateDataResponse struct {
	Sector               string  `json:"sector,omitempty"`
	Tags                 string  `json:"tags,omitempty"`
	Count                int     `json:"count,omitempty"`
	Volume               float64 `json:"volume,omitempty"`
	Marketcap            float64 `json:"marketcap,omitempty"`
	TwentyFourHourChange float64 `json:"24hr_change"`
}

// Structs for capturing messari response
type AggregateData struct {
	Data []AggregateDataItem `json:"data"`
}

type AggregateDataItem struct {
	Symbol  string `json:"symbol,omitempty"`
	Name    string `json:"name,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Metrics struct {
		MarketData struct {
			Price                float64 `json:"price_usd,omitempty"`
			Volume               float64 `json:"volume_last_24_hours,omitempty"`
			TwentyFourHourChange float64 `json:"percent_change_usd_last_24_hours,omitempty"`
			OHLCVLast24Hour      struct {
				Open float64 `json:"open"`
			} `json:"ohlcv_last_24_hour"`
		} `json:"market_data"`
		Marketcap struct {
			Marketcap float64 `json:"current_marketcap_usd,omitempty"`
		} `json:"marketcap"`
		MiscData struct {
			Sectors []string `json:"sectors,omitempty"`
			Tags    []string `json:"tags,omitempty"`
		} `json:"misc_data,omitempty"`
	} `json:"metrics,omitempty"`
}

// helper for checking values in sector/tags slice
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.ToLower(item) == strings.ToLower(val) {
			return true
		}
	}
	return false
}

// helper function to return yesterday's marketcap by undoing 24 hr % change
func adjustValueForYesterday(val float64, percentChange float64) float64 {
	var isPositive bool
	if percentChange > 0 {
		isPositive = true
	}

	// divide if change from yesterday was positive
	if isPositive {
		formattedPercentage := 1 + abs(percentChange)/100
		return val / formattedPercentage
	} else {
		formattedPercentage := 1 - abs(percentChange)/100
		return val / formattedPercentage
	}
}

// absolute value helper
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
