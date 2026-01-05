// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/logx"
	"ysgit.lunalabs.cn/products/go-common/middleware"

	"github.com/gpencil/upload/internal/config"
	"github.com/gpencil/upload/internal/handler"
	"github.com/gpencil/upload/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/uploadapi.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	var ctx *svc.ServiceContext
	conf.MustLoad(*configFile, &c)
	viper.SetConfigFile(*configFile)
	viper.ReadInConfig()
	viper.Unmarshal(&c)
	viper.OnConfigChange(func(e fsnotify.Event) {
		viper.Unmarshal(&c)
		ctx.Config = c
		logx.Infof("config file changed: %+v", c)
	})
	viper.WatchConfig()
	server := middleware.MustNewRestServer(c.RestConf)
	defer server.Stop()

	ctx = svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
