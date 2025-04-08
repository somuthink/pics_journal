package queue

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/somuthink/pics_journal/core/internal/config"
)

var RDB *redis.Client

func Initialize() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Cfg.SESSIONS_HOST, config.Cfg.SESSIONS_PORT),
		Password: config.Cfg.SESSIONS_PASSWORD,
		DB:       0,
	})
}
