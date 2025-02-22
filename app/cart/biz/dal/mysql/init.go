package mysql

import (
	"fmt"
	"os"
	"github.com/North-al/douyin-mall/app/cart/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	// 初始化数据库连接
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	fmt.Println("DSN:",dsn)//打印dsn调试
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	
	
}
