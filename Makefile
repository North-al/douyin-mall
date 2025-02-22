.PHONY: gen-rpc-client-auth
.PHONY: gen-rpc-server-auth
.PHONY: gen-rpc-client-user
.PHONY: gen-rpc-server-user
.PHONY: gen-rpc-client-product
.PHONY: gen-rpc-server-product
.PHONY: gen-rpc-client-cart
.PHONY: gen-rpc-server-cart
.PHONY: gen-cart

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

		
gen-rpc-client-product:
	cd $(RPC_GEN_DIR) && \
	cwgo client \
		--type RPC \
		--service product \
		--module $(MODULE)/rpc_gen \
		-I ../idl \
		--idl ../idl/product.proto


gen-rpc-server-product:
	cd app/product && \
	cwgo server \
		--type RPC \
		--service product \
		--module $(MODULE)/app/product \
		--pass "-use $(MODULE)/rpc_gen/kitex_gen" \
		-I ../../idl \
		--idl ../../idl/product.proto

				
gen-rpc-client-cart:
	cd $(RPC_GEN_DIR) && \
	cwgo client \
		--type RPC \
		--service cart \
		--module $(MODULE)/rpc_gen \
		-I ../idl \
		--idl ../idl/cart.proto


gen-rpc-server-cart:
	cd app/cart && \
	cwgo server \
		--type RPC \
		--service cart \
		--module $(MODULE)/app/cart \
		--pass "-use $(MODULE)/rpc_gen/kitex_gen" \
		-I ../../idl \
		--idl ../../idl/cart.proto
		
gen-cart: gen-rpc-client-cart gen-rpc-server-cart