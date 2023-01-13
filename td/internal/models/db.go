package models

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"gas-td-importer/td/internal/config"
	"gas-td-importer/td/internal/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/taosdata/driver-go/v2/taosRestful"
	//_ "github.com/taosdata/driver-go/v2/taosSql"
	_ "github.com/godror/godror"
)

var PMap = make(map[string]PointInfo) // 全局唯一线程安全PointsMap
var PLocker sync.RWMutex              // PMap读写锁

// InitDB 初始化关系库连接
func InitDB(c config.Config) *sqlx.DB {
	dsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s:%d/%s"`, c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.ServiceName)
	var err error
	db, err := sqlx.Open("godror", dsn)
	if err != nil {
		log.Fatalf("db init error:%v", err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("db init error:%v", err)
	}
	log.Println("db init success")
	return db
}

func SyncPointsMap(db *sqlx.DB, c config.Config) {
	// 周期性调用获取并保存数据
	beat := 0
	if c.DB.PointsMapBeat <= 0 {
		beat = 60
	} else {
		beat = c.DB.PointsMapBeat
	}
	ticker := time.Tick(time.Duration(beat) * time.Minute)
	for range ticker {
		log.Println("SyncPointsMap")
		InitPointsMap(db, c)
	}
}

// InitPointsMap 初始化PointsMap
func InitPointsMap(db *sqlx.DB, c config.Config) {

	// 查询所有points
	rows, err := queryPointsRows(db, c.DB.PointsTable)
	if err != nil {
		log.Fatal("InitPointsMap error：" + err.Error())
	}

	// 遍历points 写入PointsMap
	count := 0
	PLocker.Lock() // get写锁
	for rows.Next() {
		var p PointInfo
		if err := rows.StructScan(&p); err != nil {
			log.Fatalf("PointsMap iterate error：" + err.Error())
		}

		PMap[p.Point] = p
		count++
	}
	PLocker.Unlock()
	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal("PointsMap iterate error：" + err.Error())
	}
	log.Println("PointsMap's length is ", count)
	if count == 0 {
		log.Fatal("PointsMap's length is 0")
	}
}

// queryPointsRows 查询点位rows
func queryPointsRows(db *sqlx.DB, table string) (*sqlx.Rows, error) {
	sql := `
        SELECT * FROM
    ` + table
	rows, err := db.Queryx(sql)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// 数值单位转换
func ConvertUnit(p *PointInfo, g *types.Gas) {
	switch p.Unit {
	case "Nm3/d":
		p.Unit = "Nm3/h"
		g.Value = g.Value * 24
	case "E4Nm3/d":
		p.Unit = "Nm3/h"
		g.Value = g.Value * 24 * 10000
	case "kPa":
		p.Unit = "MPa"
		g.Value, _ = strconv.ParseFloat(strconv.FormatFloat(g.Value/1000, 'f', 2, 64), 64)

	}
}
