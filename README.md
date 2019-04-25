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
    |—— modules
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
    ├── routers
    ├── .env // 配置文件，自己生成
    ├── .env.example // 配置文件，模板，仅提供结构和示范
    └── tests
```

## 技术选型

- Kafka - 高吞吐量的分布式发布订阅消息系统
- GORM - ORM库 [[项目地址]](https://github.com/jinzhu/gorm) [[官方文档]](http://gorm.io/) [[中文文档]](http://gorm.book.jasperxu.com/)
- Cobra - 命令行包 [[项目地址]](https://github.com/spf13/cobra)

## 软件需求

- go 1.12

## 代理设置

- GOPROXY - A global proxy for go modules [[官网及文档]](https://goproxy.io/) ，[[备用地址]](https://goproxy.cn/)

  - mac 或 linux用户执行：`export GOPROXY=https://goproxy.io`
  - windows用户操作流程
    - **方法1** windows 用户开启powerShell必须使用管理员权限
    - **方法2** windows 用户进入 我的电脑 >> 右键 >> 属性 >> 高级系统设置 >> 环境变量 >> 系统变量 中新增一项，名字为GOPROXY 值为： https://goproxy.io
  - 如果遇到 https://goproxy.io 不能使用的情况，请使用 https://goproxy.cn 代替

## 软件本地化配置

```
cp .env.example .env
```

## 本地部署

运行
```
go build ./...
```

测试

```
go test ./...
```

**参数**

- `./...` : 配合当前模块所有相关包

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
> **NOTE**  
> `go mod tidy` 执行后，会下载相关包到 `$GOPATH/pkg/mod` 目录下

## 版本发布

> **参考**  
> @https://github.com/golang/go/wiki/Modules#how-to-prepare-for-a-release

## 如何交流、反馈、参与贡献？

xxx

## License

[MIT license](http://opensource.org/licenses/MIT)