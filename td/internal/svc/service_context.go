package svc

import (
	"database/sql"

	"gas-td-importer/td/internal/config"
	"gas-td-importer/td/internal/models"
)

type ServiceContext struct {
	Config config.Config
	Engine *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c),
	}
}
