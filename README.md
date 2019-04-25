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

## Swagger Doc Comment Example

```
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
```
