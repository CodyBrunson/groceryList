package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *Config) handlerCreateNewList(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "No name provided", err)
		return
	}

	_, err = cfg.db.CreateList(r.Context(), params.Name)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating list", err)
	}
	respondWithJSON(w, http.StatusCreated, nil)
}
