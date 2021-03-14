package aggregate

import (
	"encoding/json"
	"fmt"
	"net/http"

	api "github.com/jonathanpatrick/messari/messari_api"
)

func AggregateHandler(w http.ResponseWriter, r *http.Request) {
	tags := r.URL.Query().Get("tags")
	sector := r.URL.Query().Get("sector")
	aggregateData := &AggregateData{}
	var err error
	if tags != "" {
		err = api.GetAggregate(tags, "", aggregateData)
	} else if sector != "" {
		err = api.GetAggregate("", sector, aggregateData)
	} else {
		err = api.GetAggregate("", "", aggregateData)
	}
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error while retrieving aggregate data. err=%v", err)))
	}
	// mapAggregateResponse(*aggregateData)
	resp, err := json.Marshal(*aggregateData)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error while marshaling json. err=%v", err)))
	}

	w.Write([]byte(resp))
}

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
