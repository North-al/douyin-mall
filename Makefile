.PHONY: gen-rpc-client-auth
.PHONY: gen-rpc-server-auth
.PHONY: gen-rpc-client-user
.PHONY: gen-rpc-server-user
.PHONY: gen-rpc-client-payment
.PHONY: gen-rpc-server-payment
.PHONY: gen-rpc-client-checkout
.PHONY: gen-rpc-server-checkout

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

gen-rpc-client-user:
	cd $(RPC_GEN_DIR) && \
	cwgo client \
		--type RPC \
		--service user \
		--module $(MODULE)/rpc_gen \
		-I ../idl \
		--idl ../idl/user.proto

gen-rpc-server-user:
	cd app/user && \
	cwgo server \
		--type RPC \
		--service user \
		--module $(MODULE)/app/user \
		--pass "-use $(MODULE)/rpc_gen/kitex_gen" \
		-I ../../idl \
		--idl ../../idl/user.proto

gen-rpc-client-payment:
	cd $(RPC_GEN_DIR) && \
	cwgo client \
		--type RPC \
		--service payment \
		--module $(MODULE)/rpc_gen \
		-I ../idl \
		--idl ../idl/payment.proto

gen-rpc-server-payment:
	cd app/payment && \
	cwgo server \
		--type RPC \
		--service payment \
		--module $(MODULE)/app/payment \
		--pass "-use $(MODULE)/rpc_gen/kitex_gen" \
		-I ../../idl \
		--idl ../../idl/payment.proto


gen-rpc-client-checkout:
	cd $(RPC_GEN_DIR) && \
	cwgo client \
		--type RPC \
		--service checkout \
		--module $(MODULE)/rpc_gen \
		-I ../idl \
		--idl ../idl/checkout.proto

gen-rpc-server-checkout:
	cd app/checkout && \
	cwgo server \
		--type RPC \
		--service checkout \
		--module $(MODULE)/app/checkout \
		--pass "-use $(MODULE)/rpc_gen/kitex_gen" \
		-I ../../idl \
		--idl ../../idl/checkout.proto