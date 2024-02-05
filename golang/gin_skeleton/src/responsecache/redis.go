package responsecache

import (
	"os"
	"time"
	"fmt"
	"log"

	"github.com/chenyahui/gin-cache/persist"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Store	*persist.RedisStore
	DefaultCacheTime	time.Duration
}

func Init() *RedisCache {
	return &RedisCache{
		Store: persist.NewRedisStore(redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr: fmt.Sprintf(
				"%s:%s",
				os.Getenv("REDIS_HOST"),
				os.Getenv("REDIS_PORT"),
			),
		})),
		DefaultCacheTime: 10 * time.Second,
	}
}
