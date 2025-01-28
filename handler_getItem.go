package main

import (
	"github.com/google/uuid"
	"net/http"
)

func (cfg *Config) handlerGetAllItems(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Items []Item
	}

	items, err := cfg.db.GetAllItems(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting items", err)
	}
	if items == nil {
		emptyItems := []Item{}
		respondWithJSON(w, http.StatusOK, emptyItems)
		return
	}
	respondWithJSON(w, http.StatusOK, items)
}

func (cfg *Config) handlerGetAllItemsForList(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Items []Item
	}

	listID := r.PathValue("listID")
	if listID == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	parsedListID, err := uuid.Parse(listID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "No list id found", err)
		return
	}
	items, err := cfg.db.GetItemsForList(r.Context(), parsedListID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting items", err)
	}

	if items == nil {
		emptyItems := []Item{}
		respondWithJSON(w, http.StatusOK, emptyItems)
		return
	}
	returnItems := []Item{}
	for _, item := range items {
		returnItems = append(returnItems, Item{
			ID:     item.ID,
			Name:   item.Name,
			Amount: item.Amount,
		})
	}

	respondWithJSON(w, http.StatusOK, returnItems)
}
