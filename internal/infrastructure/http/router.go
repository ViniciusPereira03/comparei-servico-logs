package http

import (
	"comparei-servico-logs/internal/app"
	"comparei-servico-logs/internal/infrastructure/http/middleware"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func NewRouter(logService *app.LogService) *mux.Router {
	_ = godotenv.Load()

	r := mux.NewRouter()

	r.Use(middleware.APIKeyMiddleware)

	r.HandleFunc("/log", CreateLog).Methods("POST")

	return r
}
