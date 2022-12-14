package models

import (
	"database/sql"
	"fmt"
	"log"

	"gas-td-importer/td/internal/config"
	_ "github.com/taosdata/driver-go/v2/taosRestful"
	//_ "github.com/taosdata/driver-go/v2/taosSql"
)

func InitTaos(c config.Config) *sql.DB {
	//var taosUri = fmt.Sprintf("%s:%s@tcp(%s:%d)/", c.TD.User, c.TD.Password, c.TD.Host, c.TD.Port)
	//fmt.Println("taosUri:", taosUri)
	//taos, err := sql.Open("taosSql", taosUri)
	//if err != nil {
	//	log.Fatal("taos init error:%v", err)
	//	return nil
	//}

	var taosDSN = fmt.Sprintf("%s:%s@http(%s:%d)/", c.TD.User, c.TD.Password, c.TD.Host, c.TD.Port)
	taos, err := sql.Open("taosRestful", taosDSN)
	if err != nil {
		log.Fatal("taos init error:%v", err)
		return nil
	}

	taos.SetMaxOpenConns(c.TD.MaxOpenConns)
	log.Println("taos init success")
	return taos
}
