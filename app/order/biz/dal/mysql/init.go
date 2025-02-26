package mysql

import (
    "os"

    "github.com/North-al/douyin-mall/app/order/biz/model"
    "github.com/North-al/douyin-mall/app/order/conf"
    "github.com/cloudwego/kitex/pkg/klog"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var (
    DB  *gorm.DB
    err error
)

func Init() {
    // 直接使用配置文件中的 DSN
    dsn := conf.GetConf().MySQL.DSN
    DB, err = gorm.Open(mysql.Open(dsn),
        &gorm.Config{
            PrepareStmt:            true,
            SkipDefaultTransaction: true,
        },
    )
    if err != nil {
        panic(err)
    }
    if os.Getenv("GO_ENV") != "online" {
        if err := DB.AutoMigrate(
            &model.Order{},
            &model.OrderItem{},
        ); err != nil {
            klog.Error(err)
        }
    }
}