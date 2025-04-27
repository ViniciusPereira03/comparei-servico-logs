package http

import (
	"comparei-servico-logs/internal/app"
	"comparei-servico-logs/internal/infrastructure/http/dto"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
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

func validaToken(w http.ResponseWriter, r *http.Request) (string, error) {
	secret := os.Getenv("USER_JWT_SECRET")

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("Missing token")
	}

	// Remover o prefixo "Bearer " se existir
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Verificar o token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}

	// Acessar os dados (claims)
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id := claims["id"]
		return fmt.Sprintf("%v", id), nil
	}

	return "", fmt.Errorf("Erro ao decodificar token")
}

func CreateLog(w http.ResponseWriter, r *http.Request) {
	userID, err_token := validaToken(w, r)
	if err_token != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err_token, "Erro ao refistrar log")
		return
	}

	var logDTO dto.CreateLogDTO
	if err := json.NewDecoder(r.Body).Decode(&logDTO); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, err, "JSON inválido")
		return
	}

	log := logDTO.ParseToLog(userID)
	err := service.CreateLog(log)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err, "Erro ao refistrar log")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Log registrado com sucesso!")
}
