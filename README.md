# douyin-mall

基于 Kitex 框架的抖音商城微服务项目。

## 项目结构

```
douyin-mall
├── app     # 微服务应用目录
│ └── auth  # 认证服务
│ └── ...   # 其他服务
├── idl     # 接口定义文件目录
└── rpc_gen # RPC 生成代码目录
    ├── kitex_gen # Kitex 生成的代码
    └── rpc       # RPC 相关代码
```

## 开发环境

该项目已经配置了devcontainer，可以直接使用的Dev Container来开发。

## 代码生成
参考项目文件的`Makefile`文件，执行`make <target>`命令来生成代码。
```bash
make <target>
```

