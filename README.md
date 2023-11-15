# learning

## 环境说明
* 语言:golang:1.21.4
* 脚手架:go-zero
* 数据库:mysql5.7

## 工程目录
```
｜——  command 服务启动配置
    ｜——  root.go 服务启动入口
    ｜——  server.go go-zero启动函数
    ｜——  stepup.go 连接入口
    ｜——  version.go 版本信息
    ｜——  apigateway 路由配置
｜——  doc 设计文件路径
｜——  migration 迁移文件路径
｜——  internal 代码工程目录
    ｜—— core 领域驱动层
    ｜—— handler 视图层
    ｜—— middleware 中间件
    ｜—— service 视图层
｜——  pkg 存放第三方工具
    ｜—— libs 封装依赖库
    ｜—— utils 工具类
｜——  README 工程说明文件
｜——  Dockerfile
｜—— .gitignore git的忽略文件
｜—— .golangci.yml golang代码检查配置
```

## 启动顺序

## 插件
1. 代码规范插件：  
* goformat 代码格式化
* golangci-lint 代码检查
* goimport 自动导包

2. 画图设计插件 
* diagrams.net

## 常用命令
* 设计文件
```

```
* 数据库迁移命令
```azure

```
* go-zero常用命令
1. api生成
```
goctl api -o test.api
```
2. 接口文档生成
```
生成到gen文件夹中
goctl api plugin -plugin goctl-swagger="swagger -filename test.json" -api test.api -dir ./gen/
redoc-cli bundle ./gen/test.json -o ./gen/apidoc.html
```
3. 生成请求
```
goctl api go --api test.api -dir gen
rm -rf *.go
rm -rf etc
rm -rf internal/config
rm -rf internal/svc
```
4. rpc接口生成
```api

```