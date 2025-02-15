package main

import (
	"net"
	"time"

	"github.com/North-al/douyin-mall/app/auth/biz/dal"
	"github.com/North-al/douyin-mall/app/auth/conf"
	"github.com/North-al/douyin-mall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kr/pretty"

	"github.com/cloudwego/kitex/pkg/registry"
	consul "github.com/kitex-contrib/registry-consul"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	dal.Init()
	opts := kitexInit()

	// 注册到 Consul
	r, err := registerToConsul()
	if err != nil {
		klog.Fatalf("Failed to register to Consul: %v", err)
		return
	}
	opts = append(opts, server.WithRegistry(r))

	// 创建并启动服务
	svr := authservice.NewServer(new(AuthServiceImpl), opts...)

	pretty.Printf("Starting auth service...\n")
	if err := svr.Run(); err != nil {
		klog.Fatalf("Service stopped with error: %v", err)
		return
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}

func registerToConsul() (registry.Registry, error) {
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	pretty.Printf("Register to Consul: %+v\n", r)
	if err != nil {
		klog.Errorf("Failed to register to Consul: %v", err)
		return nil, err
	}
	return r, nil
}
