package main

import (
	"github.com/google/uuid"
	"net/http"
)

func (cfg *Config) handlerDeleteListByID(w http.ResponseWriter, r *http.Request) {
	listID := r.PathValue("listID")
	if listID == "" {
		respondWithError(w, http.StatusNotFound, "List ID not found", nil)
		return
	}
	parsedID, err := uuid.Parse(listID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "List ID not found", err)
		return
	}

	list, err := cfg.db.GetListByID(r.Context(), parsedID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "List ID not found", err)
		return
	}

	err = cfg.db.DeleteListByID(r.Context(), list.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "List ID not found", err)
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}
