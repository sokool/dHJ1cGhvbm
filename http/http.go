package http

import (
	"github.com/gorilla/mux"
	"net/http"
	"truphone/model"
)

func New(address string, r model.Devices) error {
	d := newDevices(r)
	m := mux.NewRouter()
	m.HandleFunc("/devices", d.add).Methods(http.MethodPost)
	m.HandleFunc("/devices", d.search).Methods(http.MethodGet)
	m.HandleFunc("/devices/{device}", d.get).Methods(http.MethodGet)
	m.HandleFunc("/devices/{device}", d.update).Methods(http.MethodPut)
	m.HandleFunc("/devices/{device}", d.delete).Methods(http.MethodDelete)

	return http.ListenAndServe(address, m)
}
