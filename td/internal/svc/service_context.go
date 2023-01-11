package svc

import (
	"database/sql"

	"gas-td-importer/td/internal/config"
	"gas-td-importer/td/internal/models"
	"github.com/jmoiron/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	TaosEngine *sql.DB
	DBEngine   *sqlx.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	taos := models.InitTaos(c)
	db := models.InitDB(c)
	models.InitPointsMap(db, c)

	// 定期更新PointsMap
	go models.SyncPointsMap(db, c)

	return &ServiceContext{
		Config:     c,
		TaosEngine: taos,
		DBEngine:   db,
	}
}
