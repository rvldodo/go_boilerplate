package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/rvldodo/boilerplate/lib/log"
)

type ResponseFormat struct {
	Status     string      `json:"status,omitempty"`
	StatusCode int         `json:"status_code,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ParseJSON(r *http.Request, v interface{}) error {
	if r.Body == nil {
		log.Error("Missing request body")
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	statusRes := true
	if status >= 400 {
		statusRes = false
	}

	return json.NewEncoder(w).Encode(&ResponseFormat{
		Status:     strconv.FormatBool(statusRes),
		StatusCode: status,
		Data:       v,
	})
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]interface{}{"error": err.Error()})
}
