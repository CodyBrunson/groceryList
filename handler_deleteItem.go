package main

import (
	"github.com/google/uuid"
	"net/http"
)

func (cfg *Config) handlerDeleteItemByID(w http.ResponseWriter, r *http.Request) {
	itemID := r.PathValue("itemID")
	if itemID == "" {
		respondWithError(w, http.StatusNotFound, "Item ID not found", nil)
		return
	}
	parsedID, err := uuid.Parse(itemID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Item ID not found", err)
		return
	}

	item, err := cfg.db.GetItemByID(r.Context(), parsedID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Item ID not found", err)
		return
	}

	err = cfg.db.DeleteItem(r.Context(), item.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Item ID not found", err)
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}
