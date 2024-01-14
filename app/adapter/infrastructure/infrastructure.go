package infrastructure

import (
	"github.com/yuorei/video-server/app/application/port"
	"github.com/yuorei/video-server/app/driver/db"
)

type Infrastructure struct {
	port.VideoRepository
	db *db.DB
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{
		db: db.NewMongoDB(),
	}
}
