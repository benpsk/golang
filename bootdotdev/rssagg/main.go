package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/benpsk/golang/bootdotdev/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not found in the environment")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("Database URL env is not found in the environment")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()
	v1Router.Post("/users", apiCfg.handleUsersCreate)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleUserGet))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handleFeedCreate))
	v1Router.Get("/feeds", apiCfg.handleGetFeeds)

	v1Router.Get("/healthz", handleReadiness)
	v1Router.Get("/err", handleErr)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server starting on port %v", port)
	log.Fatal(srv.ListenAndServe())
}
