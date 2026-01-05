// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"log"
	"time"

	"github.com/gpencil/upload/internal/config"
	"github.com/gpencil/upload/internal/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Query  *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	db, err := gorm.Open(mysql.Open(c.DB.DataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB: %v", err)
	}

	sqlDB.SetMaxOpenConns(c.DB.MaxOpenConns)
	sqlDB.SetMaxIdleConns(c.DB.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.DB.ConnMaxLifetime) * time.Second)

	return &ServiceContext{
		Config: c,
		DB:     db,
		Query:  query.Use(db),
	}
}
