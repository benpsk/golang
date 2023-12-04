package main

import (
	"net/http"

	"github.com/benpsk/golang/bootdotdev/rssagg/internal/auth"
	"github.com/benpsk/golang/bootdotdev/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Could'nt find api key")
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Could'nt get user")
			return
		}
		handler(w, r, user)
	}
}
