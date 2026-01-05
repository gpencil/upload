package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:12345678@tcp(localhost:3306)/tts_dev?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect database: %v\n", err)
		os.Exit(1)
	}

	// 创建生成器
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./internal/dal/query", // 查询代码输出目录
		ModelPkgPath:  "./internal/dal/model", // Model输出目录
		FieldNullable: true,                   // 字段可为空
	})

	g.UseDB(db)

	// 生成各个表的Model
	g.ApplyBasic(
		g.GenerateModelAs("voices", "Voices"),
		g.GenerateModelAs("scenarios", "Scenarios"),
	)

	g.Execute()
}
