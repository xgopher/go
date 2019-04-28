# 示例模块

示例模块，请根据实际需求删减

## 结构说明

```
.
├── cmd             // 命令行程序
├── controllers     // 控制器
├── services        // 服务层(辅助controller，如：发邮件，短信服务...)
├── models          // 模型层
├── repositories    // 仓储层(辅助model)
├── views           // 视图层(可选)
├── utils           // 工具类(存放辅助函数，如：utf8转码...)
├── conf            // 配置文件(可选)
└── module.go       // 模块启动引导文件
```

## RESTFUL 示例

|方法        |路径                       |动作     |
|----        |----                      |----     |
|GET         | /demo/objects          | index   |
|POST        | /demo/objects          | store   |
|GET         | /demo/objects/{object} | show    |
|PUT/PATCH   | /demo/objects/{object} | update  |
|DELETE      | /demo/objects/{object} | destroy |