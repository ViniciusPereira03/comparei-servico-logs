package publisher

import (
	"comparei-servico-logs/config"
	"comparei-servico-logs/internal/domain/user"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	config.LoadConfig()
	host := fmt.Sprintf("%v:%v", os.Getenv("REDIS_MESSAGING_HOST"), os.Getenv("REDIS_MESSAGING_PORT"))
	rdb = redis.NewClient(&redis.Options{
		Addr: host,
	})
}

func PubUpdateLevelUser(u *user.User) error {
	ctx := context.Background()

	payload, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return fmt.Errorf("erro ao codificar payload: %v", err)
	}
	_, err = rdb.Publish(ctx, "update_level_user", string(payload)).Result()
	if err != nil {
		return fmt.Errorf("erro ao publicar mensagem no Redis: %v", err)
	}

	return nil
}
