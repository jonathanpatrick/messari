package asset

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response, _ := json.Marshal(vars)

	w.Write([]byte(response))
}
