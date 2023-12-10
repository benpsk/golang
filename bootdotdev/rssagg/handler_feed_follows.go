package main

import (
	"net/http"

	"github.com/benpsk/golang/bootdotdev/rssagg/internal/database"
)

func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
}