# Gin-Learning

[Source Article](https://studygolang.com/subject/194)


## Folder Structure

```js
.
├── README.md
├── conf
│   └── app.ini
├── main.go
├── middlewares
├── models
│   ├── article.go
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── cerror
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       └── pagination.go
├── resources
│   ├── api.v1
│   └── db.sql
├── routers
│   ├── api
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
└── runtime
```

- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime 应用运行时数据
- resources 一些资源文件
