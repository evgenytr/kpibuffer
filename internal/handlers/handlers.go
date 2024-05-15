// Package handlers contains handlers for all http requests to buffer server
package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/evgenytr/kpibuffer/internal/domain"
	"github.com/evgenytr/kpibuffer/internal/interfaces"
)

// StorageHandler struct contains Storage interface.
type StorageHandler struct {
	storage interfaces.Storage
}

// NewStorageHandler returns StorageHandler for passed interface.
func NewStorageHandler(storage interfaces.Storage) *StorageHandler {
	return &StorageHandler{
		storage: storage,
	}
}

func (h *StorageHandler) GetFactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	totalFactsInStorage := h.storage.Len()
	fmt.Fprintf(w, "Total facts stored %v", totalFactsInStorage)
	w.WriteHeader(http.StatusOK)
}

func (h *StorageHandler) PostFactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	var currFact domain.Fact
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)

	if err != nil {
		processBadRequest(w, err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &currFact)

	if err != nil {
		processBadRequest(w, err)
		return
	}

	h.storage.PushFact(&currFact)

	w.WriteHeader(http.StatusOK)
}

func (h *StorageHandler) PostFactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	var currFacts []domain.Fact
	var buf bytes.Buffer

	_, err := buf.ReadFrom(r.Body)

	if err != nil {
		processBadRequest(w, err)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &currFacts)

	if err != nil {
		processBadRequest(w, err)
		return
	}

	for _, currFact := range currFacts {
		h.storage.PushFact(&currFact)
	}

	w.WriteHeader(http.StatusOK)
}

func processBadRequest(res http.ResponseWriter, err error) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusBadRequest)

	_, errOut := fmt.Fprintf(res, "Bad request, error %v", err)
	if errOut != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
}
