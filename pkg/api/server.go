package api

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status bool   `json:Status`
	Data   string `json:"Data"`
}

// NewServer creates a new http server instance with given mux
func NewServer(mux *http.ServeMux) *http.Server {
	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	return server
}

func respond(res http.ResponseWriter, data string, status bool, code int) {
	res.WriteHeader(code)
	rsp := response{
		Data:   data,
		Status: status,
	}
	json.NewEncoder(res).Encode(rsp)
}
