package main

import (
	"encoding/json"
	"github.com/CodyBrunson/groceryList/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type Item struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Amount string    `json:"amount"`
}

func (cfg *Config) handlerAddItem(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name   string `json:"name"`
		Amount string `json:"amount"`
		ListID string `json:"listID"`
	}

	type response struct {
		Item
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request", err)
		return
	}
	parsedListID, err := uuid.Parse(params.ListID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid list id", err)
		return
	}
	newItem := database.CreateItemParams{
		Name:   params.Name,
		Amount: params.Amount,
		ListID: parsedListID,
	}
	_, err = cfg.db.CreateItem(r.Context(), newItem)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating item", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, nil)

}
