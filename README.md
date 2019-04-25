## 项目说明

go 是一个基于 `beego` 框架轻量级快速开发平台

## 项目结构

```
.
└── go
    ├── LICENSE
    ├── README.md
    ├── conf
    ├── controllers
    ├── docs
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
    ├── routers
    └── tests
```

## 技术选型

- Kafka # 高吞吐量的分布式发布订阅消息系统

## 软件需求

- go 1.12

## 本地部署

```
go mod tidy
```

执行后，会下载相关包到 `$GOPATH/pkg/mod` 目录下

**命令说明**

```
go mod init:初始化modules
go mod download:下载modules到本地cache
go mod edit:编辑go.mod文件，选项有-json、-require和-exclude，可以使用帮助go help mod edit
go mod graph:以文本模式打印模块需求图
go mod tidy:检查，删除错误或者不使用的modules，下载没download的package
go mod vendor:生成vendor目录
go mod verify:验证依赖是否正确
go mod why：查找依赖

go test    执行一下，自动导包

go list -m  主模块的打印路径
go list -m -f={{.Dir}}  print主模块的根目录
go list -m all  查看当前的依赖和版本信息
```
## 如何交流、反馈、参与贡献？

xxx

## License

[MIT license](http://opensource.org/licenses/MIT)