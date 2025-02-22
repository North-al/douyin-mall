package mysql

import (
	"fmt"

	"github.com/North-al/douyin-mall/app/user/conf"
	"github.com/North-al/douyin-mall/app/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	fmt.Println(conf.GetConf().MySQL.DSN)
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// 自动迁移
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
