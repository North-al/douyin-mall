# douyin-mall

字节青训营，基于 Kitex 框架的抖音商城微服务项目。项目地址：[https://github.com/North-al/douyin-mall](https://github.com/North-al/douyin-mall)

## 项目结构

```
douyin-mall
├── app               # 微服务应用目录
│   ├── auth          # 认证服务
│   │   ├── conf      # 配置文件
│   │   ├── biz       # 业务逻辑
│   │   └── main.go   # 服务入口
│   ├── user          # 用户服务
│   │   ├── model     # 数据模型
│   │   ├── biz       # 业务逻辑
│   │   │   └── query # 数据库操作封装
│   │   └── dal       # 数据访问层
│   └── ...           # 其他服务
├── idl               # 接口定义文件目录
└── rpc_gen           # RPC 生成代码目录
    ├── kitex_gen     # Kitex 生成的代码
    └── rpc           # RPC 相关代码
```

## 开发环境

该项目已经配置了 Dev Container，包含以下服务：

- MySQL 8.0
- Redis latest
- Consul (服务发现)

### 使用 Dev Container 开发

1. 安装 VS Code 和 Dev Container 插件
2. 克隆项目并在 VS Code 中打开
3. 点击左下角绿色按钮，选择 "Reopen in Container"

### 服务端口

- MySQL: 3316
- Redis: 6389
- Consul UI: http://localhost:8500

## 代码生成

参考项目文件的 `Makefile` 文件，执行 `make <target>` 命令来生成代码。

```bash
# 生成特定服务的代码
make gen-rpc-client-auth # 生成认证服务客户端代码
make gen-rpc-server-auth # 生成认证服务服务端代码
make gen-rpc-client-user # 生成用户服务客户端代码
make gen-rpc-server-user # 生成用户服务服务端代码
```

## 服务说明

### 认证服务 (Auth)
- Token 的签发和验证
  - DeliverTokenByRPC: 根据用户ID签发 token
  - VerifyTokenByRPC: 验证 token 的有效性

### 用户服务 (User)
- 用户账户管理
  - Register: 用户注册（邮箱、密码）
  - Login: 用户登录（邮箱、密码）

## 配置文件

配置文件位于各服务的 `conf` 目录下，按环境区分：

```
app/*/conf/
├── test/        # 测试环境配置
├── dev/         # 开发环境配置
└── online/      # 生产环境配置
```

## 开发指南

1. 所有新服务都应该注册到 Consul 进行服务发现
2. 使用 GORM 进行数据库操作
3. 遵循项目的目录结构和命名规范
4. 确保添加适当的日志和错误处理

## 参考资料

- [Kitex 官方文档](https://www.cloudwego.io/docs/kitex/)
- [Consul 官方文档](https://www.consul.io/docs)
- [GORM 官方文档](https://gorm.io/docs)
