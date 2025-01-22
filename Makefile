.PHONY: gen-rpc-client-auth
.PHONY: gen-rpc-server-auth

# 定义变量以便于维护和修改
RPC_GEN_DIR := rpc_gen
MODULE := github.com/North-al/douyin-mall

# 生成 RPC 客户端代码
gen-rpc-client-auth:
	cd $(RPC_GEN_DIR) && \
	cwgo client \
		--type RPC \
		--service auth \
		--module $(MODULE)/rpc_gen \
		-I ../idl \
		--idl ../idl/auth.proto

# 生成 RPC 服务端代码
gen-rpc-server-auth:
	cd app/auth && \
	cwgo server \
		--type RPC \
		--service auth \
		--module $(MODULE)/app/auth \
		--pass "-use $(MODULE)/rpc_gen/kitex_gen" \
		-I ../../idl \
		--idl ../../idl/auth.proto
