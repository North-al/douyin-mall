package mysql

import (
	"github.com/North-al/douyin-mall/app/payment/biz/model"
	"github.com/North-al/douyin-mall/app/payment/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {

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
	err = DB.AutoMigrate(&model.PaymentLog{})
	if err != nil {
		panic(err)
	}
}
