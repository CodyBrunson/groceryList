package main

import "net/http"

func (cfg *Config) handlerGetAllLists(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	lists, err := cfg.db.GetAllLists(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting all lists", err)
		return
	}

	if lists == nil {
		emptyLists := []response{}
		respondWithJSON(w, http.StatusOK, emptyLists)
		return
	}

	respondWithJSON(w, http.StatusOK, lists)
}
