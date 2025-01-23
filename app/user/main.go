package main

import (
	"net"
	"time"

	"github.com/North-al/douyin-mall/app/user/biz/dal"
	"github.com/North-al/douyin-mall/app/user/biz/rpc"
	"github.com/North-al/douyin-mall/app/user/conf"
	"github.com/North-al/douyin-mall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/kr/pretty"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	opts := kitexInit()
	dal.Init()
	// 注册到 Consul
	r, err := registerToConsul()
	if err != nil {
		klog.Fatalf("Failed to register to Consul: %v", err)
		return
	}
	opts = append(opts, server.WithRegistry(r))

	rpc.InitAuthClient(&rpc.AuthClientConfig{
		ServiceName: "auth",
		ConsulAddr:  conf.GetConf().Registry.RegistryAddress[0],
	})

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err = svr.Run()
	if err != nil {
		klog.Error(err.Error())
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
