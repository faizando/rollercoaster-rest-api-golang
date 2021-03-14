package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Coaster struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	InPark       string `json:"inPark"`
	Height       int    `json:"height"`
}

type coasterHandlers struct {
	sync.Mutex
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

	// lock store to disallow concurrent operations
	h.Lock()
	i := 0
	for _, c := range h.store {
		coasters[i] = c
		i++
	}
	h.Unlock()

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
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // cannot read body
		w.Write([]byte(err.Error()))
		return
	}

	// check for content type
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType) // if send other than json to server
		w.Write([]byte(fmt.Sprintf("need content type 'application/json' but got '%s'", ct)))
		return
	}

	// unmarshal body data
	var coaster Coaster
	err2 := json.Unmarshal(bodyBytes, &coaster)
	if err2 != nil {
		w.WriteHeader(http.StatusBadRequest) // cannot umarshall the data sent to server
		w.Write([]byte(err2.Error()))
		return
	}

	coaster.ID = fmt.Sprintf("%d", time.Now().UnixNano())

	h.Lock()
	h.store[coaster.ID] = coaster
	defer h.Unlock()
}

func newCoasterHandler() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{},
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
