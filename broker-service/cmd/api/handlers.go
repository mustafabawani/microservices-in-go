package main

import (
	"net/http"
)

type jsonReponse struct{
	Error bool `json:"error"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}
func (app *Config) Broker(w http.ResponseWriter, r *http.Request){
	payload := jsonReponse{
		Error: false,
		Message: "hit the broker",
	}

	_ = app.writeJSON(w,http.StatusOK,payload)
	
}