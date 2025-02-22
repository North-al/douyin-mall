package rpc

import (
	"sync"

	"github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/kr/pretty"
)

var (
	authClient authservice.Client
	once       sync.Once
)

type AuthClientConfig struct {
	ServiceName string
	ConsulAddr  string // Consul 地址
}

func initAuthClient(cfg *AuthClientConfig) {
	resolver, err := consul.NewConsulResolver(cfg.ConsulAddr)
	if err != nil {
		pretty.Printf("创建 Consul 解析器失败: %v", err)
		panic(err)
	}

	opts := []client.Option{
		// 使用 Consul 进行服务发现
		client.WithResolver(resolver),
	}
	authClient, err = authservice.NewClient(cfg.ServiceName, opts...)
	if err != nil {
		pretty.Printf("初始化 Auth RPC 客户端失败: %v", err)
		panic(err)
	}

}

func InitAuthClient(cfg *AuthClientConfig) {
	once.Do(func() {
		initAuthClient(cfg)
	})
}

func GetAuthClient() authservice.Client {
	if authClient == nil {
		panic("Auth RPC 客户端未初始化")
	}
	return authClient
}
