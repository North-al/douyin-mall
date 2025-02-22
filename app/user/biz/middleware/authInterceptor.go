package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/North-al/douyin-mall/app/user/biz/service"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

// AuthMiddleware is a Kitex middleware for token authentication
func AuthMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		ri := rpcinfo.GetRPCInfo(ctx)
		method := ri.To().Method()

		excludedMethods := map[string]bool{
			"Register": true,
			"Login":    true,
		}

		if excludedMethods[method] {
			return next(ctx, req, resp)
		}

		// 从元数据中获取 Authorization 头部信息
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return errors.New("missing metadata")
		}

		fmt.Println("md", md)

		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return errors.New("authorization header missing")
		}

		parts := strings.Split(authHeader[0], " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return errors.New("authorization header format must be Bearer {token}")
		}

		token := parts[1]
		userId, err := service.VerifyToken(ctx, token)
		if err != nil {
			return errors.New("invalid token: " + err.Error())
		}

		// Add userId to the context
		ctx = context.WithValue(ctx, "userId", userId)
		return next(ctx, req, resp)
	}
}
