package aggregate

import (
	"encoding/json"

	"net/http"
)

func AggregateHandler(w http.ResponseWriter, r *http.Request) {
	var response []byte
	tags := r.URL.Query().Get("tags")
	sector := r.URL.Query().Get("sector")
	if tags != "" {
		response, _ = json.Marshal(tags)
	} else if sector != "" {
		response, _ = json.Marshal(sector)
	}

	w.Write([]byte(response))
}
