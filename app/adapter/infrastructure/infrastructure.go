package infrastructure

import (
	"github.com/redis/go-redis/v9"
	"github.com/yuorei/video-server/app/driver/db"
	r "github.com/yuorei/video-server/app/driver/redis"
)

type Infrastructure struct {
	db    *db.DB
	redis *redis.Client
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		db:    db.NewMySQLDB(),
		redis: r.ConnectRedis(),
	}
}
