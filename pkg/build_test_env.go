package pkg

import (
	"fmt"
	"os"
)

type BuildTestEnv struct {
	Env string
}

func NewBuildTestEnv(env string) *BuildTestEnv {
	return &BuildTestEnv{Env: env}
}

func (b *BuildTestEnv) SetTestEnv() {
	// 设置测试环境
	err := os.Setenv("GO_ENV", b.Env)
	if err != nil {
		panic(fmt.Sprintf("设置环境变量失败: %v", err))
	}
}

// SetWorkDir 设置工作目录，如果未提供路径则默认使用 "../../"
func (b *BuildTestEnv) SetWorkDir(path ...string) {
	defaultPath := "../../"
	if len(path) > 0 {
		defaultPath = path[0]
	}

	// 切换工作目录
	err := os.Chdir(defaultPath)
	if err != nil {
		panic(fmt.Sprintf("切换工作目录失败: %v", err))
	}
}
