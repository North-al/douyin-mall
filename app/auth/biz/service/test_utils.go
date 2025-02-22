package service

import (
	"github.com/North-al/douyin-mall/app/auth/biz/dal"
	"github.com/North-al/douyin-mall/pkg"
)

func InitTestEnv() {
	build := pkg.NewBuildTestEnv("test")
	build.SetTestEnv()
	build.SetWorkDir()
	dal.Init()
}
