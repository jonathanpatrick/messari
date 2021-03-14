package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanpatrick/messari/aggregate"
	"github.com/jonathanpatrick/messari/asset"
	api "github.com/jonathanpatrick/messari/messari_api"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/aggregate", AggregateHandler)
	r.HandleFunc("/api/asset/{asset}", AssetHandler)
	http.Handle("/", r)
	http.ListenAndServe(":80", r)
}

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	assetData := &asset.AssetData{}
	err := api.GetAsset(vars["asset"], assetData)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error while retrieving data. err=%v", err)))
	}

	resp, err := json.Marshal(asset.MapAssetResponse(*assetData))
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error while marshaling json. err=%v", err)))
	}

	w.Write([]byte(resp))
}

func AggregateHandler(w http.ResponseWriter, r *http.Request) {
	tags := r.URL.Query().Get("tags")
	sector := r.URL.Query().Get("sector")
	aggregateData := &aggregate.AggregateData{}
	err := api.GetAggregate(aggregateData)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error while retrieving aggregate data. err=%v", err)))
		return
	}

	aggregateOutput, err := aggregate.ProcessAggregateResponse(tags, sector, *aggregateData)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error while parsing aggregate data. err=%v", err)))
		return
	}
	resp, err := json.Marshal(aggregateOutput)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error while marshaling json. err=%v", err)))
		return
	}

	w.Write([]byte(resp))
}
