package subscriber

import (
	"comparei-servico-logs/config"
	"comparei-servico-logs/internal/app"
	"comparei-servico-logs/internal/domain/promer"
	"comparei-servico-logs/internal/domain/user"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var log_service *app.LogService
var user_service *app.UserService

func init() {
	config.LoadConfig()
	host := fmt.Sprintf("%v:%v", os.Getenv("REDIS_MESSAGING_HOST"), os.Getenv("REDIS_MESSAGING_PORT"))
	rdb = redis.NewClient(&redis.Options{
		Addr: host,
	})
}

// Função para injetar o user_service
func SetUserService(service *app.UserService, log *app.LogService) {
	user_service = service
	log_service = log
}

func Run() {
	go subCreateUser()
	go subNewProduct()
	go subUpdateProduct()
}

func subCreateUser() error {
	ctx := context.Background()

	sub := rdb.Subscribe(ctx, "user_created")
	ch := sub.Channel()

	for msg := range ch {
		var user user.User
		err := json.Unmarshal([]byte(msg.Payload), &user)
		if err != nil {
			fmt.Println("[ERRO] Erro ao decodificar payload de mensageria:", err)
			continue
		}

		err_create := user_service.CreateUser(&user)
		if err_create != nil {
			fmt.Println("[ERRO] Erro ao criar user nos logs:", err_create)
		}
	}

	return nil
}

func subNewProduct() error {
	ctx := context.Background()

	sub := rdb.Subscribe(ctx, "new_product")
	ch := sub.Channel()

	for msg := range ch {
		var promer promer.Promer
		err := json.Unmarshal([]byte(msg.Payload), &promer)
		if err != nil {
			fmt.Println("[ERRO] Erro ao decodificar payload de mensageria:", err)
			continue
		}

		err = log_service.CreateLog(promer.ParseToCreateLog())
	}

	return nil
}

func subUpdateProduct() error {
	ctx := context.Background()

	sub := rdb.Subscribe(ctx, "update_product")
	ch := sub.Channel()

	for msg := range ch {
		var promer promer.Promer
		err := json.Unmarshal([]byte(msg.Payload), &promer)
		if err != nil {
			fmt.Println("[ERRO] Erro ao decodificar payload de mensageria:", err)
			continue
		}

		err = log_service.CreateLog(promer.ParseToUpdateLog())
	}

	return nil
}
