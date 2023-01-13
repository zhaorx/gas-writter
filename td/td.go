package main

import (
	"flag"
	"fmt"

	"gas-td-importer/td/internal/config"
	"gas-td-importer/td/internal/handler"
	"gas-td-importer/td/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/td-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 日志
	logx.DisableStat()
	logx.MustSetup(c.Log)
	//logx.CollectSysLog()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer ctx.DBEngine.Close() // 程序退出时关闭库连接
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
