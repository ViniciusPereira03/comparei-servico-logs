package http

import (
	"comparei-servico-logs/internal/app"
	"comparei-servico-logs/internal/infrastructure/http/dto"
	"encoding/json"
	"net/http"
)

var service *app.LogService

func IniHandlers(logService *app.LogService) {
	service = logService
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, err error, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error":    err.Error(),
		"mensagem": message,
	})
}

func CreateLog(w http.ResponseWriter, r *http.Request) {
	var logDTO dto.CreateLogDTO
	if err := json.NewDecoder(r.Body).Decode(&logDTO); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err, "JSON inválido")
		return
	}

	log := logDTO.ParseToLog()
	err := service.CreateLog(log)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err, "Erro ao refistrar log")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(log)
}
