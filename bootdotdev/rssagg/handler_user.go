package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/benpsk/golang/bootdotdev/rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handleUsersCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could't decode paramters")
		return
	}
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coun't create user")
		return
	}
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
