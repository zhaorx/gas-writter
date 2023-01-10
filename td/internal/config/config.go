package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	TD struct {
		Host         string
		Port         int
		User         string
		Password     string
		MaxOpenConns int
		DataBase     string
		STable       string
	}

	DB struct {
		PointsTable   string
		PointsMapBeat int
		Host          string
		Port          int
		ServiceName   string
		Username      string
		Password      string
	}
}
