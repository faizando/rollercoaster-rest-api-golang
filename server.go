package main

import (
	"encoding/json"
	"net/http"
)

type Coaster struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	InPark       string `json:"inPark"`
	Height       int    `json:"height"`
}

type coasterHandlers struct {
	store map[string]Coaster
}

func (h *coasterHandlers) coasters(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))

	i := 0
	for _, c := range h.store {
		coasters[i] = c
		i++
	}
	jsonBytes, err := json.Marshal(coasters)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
func (h *coasterHandlers) post(w http.ResponseWriter, r *http.Request) {
}

func newCoasterHandler() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1": Coaster{
				ID:           "id1",
				Name:         "Furry",
				Manufacturer: "Adidas",
				InPark:       "ThorpePark",
				Height:       99,
			},
		},
	}
}
func main() {
	coasterHandlers := newCoasterHandler()

	http.HandleFunc("/coasters", coasterHandlers.coasters)

	err := http.ListenAndServe("localhost:8090", nil)
	if err != nil {
		panic(err)
	}
}
