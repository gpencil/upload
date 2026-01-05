// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"ysgit.lunalabs.cn/products/go-common/middleware"
)

type Config struct {
	rest.RestConf
	DB     middleware.GormConfig
	Aliyun AliyunConfig
}

type MysqlConfig struct {
	DataSource string
}

type AliyunConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	BucketName      string
	IconDir         string
	PreviewDir      string
}
