package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanpatrick/messari/aggregate"
	"github.com/jonathanpatrick/messari/asset"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/aggregate", aggregate.AggregateHandler)
	r.HandleFunc("/api/asset/{asset}", asset.AssetHandler)
	http.Handle("/", r)
	http.ListenAndServe(":80", r)
}
