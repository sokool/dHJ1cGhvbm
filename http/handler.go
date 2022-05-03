package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"io"
	"net/http"
)

type handler struct{}

func (d handler) write(w http.ResponseWriter, payload any) {
	if err, ok := payload.(error); ok && err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	if payload == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d handler) read(w http.ResponseWriter, r *http.Request, payload any) bool {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil && err != io.EOF {
		d.write(w, err)
		return false
	}

	return true
}

func (d handler) url(w http.ResponseWriter, r *http.Request, payload any) bool {
	if err := schema.NewDecoder().Decode(payload, r.URL.Query()); err != nil {
		d.write(w, err)
		return false
	}

	return true
}

func (d handler) param(r *http.Request, name string) string {
	if p, ok := mux.Vars(r)[name]; ok {
		return p
	}

	return ""
}
