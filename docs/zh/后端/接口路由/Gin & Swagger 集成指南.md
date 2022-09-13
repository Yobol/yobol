# Gin 集成 Swagger 自动生成 RESTful 文档

> [gin-swagger](https://github.com/swaggo/gin-swagger) 中间件用于为 `Gin` 自动生成 `Swagger2.0` 的 `RESTful` 文档。

## 声明式注释

`gin-swagger` 提供[声明式注释](https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format)帮助开发者自动为接口生成 `Swagger2.0` 的 `RESTful` 文档。 

### [API 概要信息](https://github.com/swaggo/swag#general-api-info)

参考官方示例：[celler/main.go]()。

`注：该部分注释需要添加到 `main.go` 文件中。`

## 下载使用

### 下载 `swag`

```shell
# 1.15 or older
$ go get -u github.com/swaggo/swag/cmd/swag

# 1.16 or newer
$ go install github.com/swaggo/swag/cmd/swag@latest
```

### 使用 `swag`

```shell
# run swag init in the server project's root folder which contains main.go file
# it will parse all comments in project and generate automatically the required files(docs folder and docs/docs.go file)
$ swag init
```

`注：记得将 $GOPATH/bin 目录加到操作系统 PATH 中。`

```shell
# import gin-swagger depedencies
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
```

在代码中引入上述依赖：

```go
import (
    "github.com/swaggo/gin-swagger" // gin-swagger middleware
    "github.com/swaggo/files"       // swagger embed files
)

func xxx() {
    route := gin.Default()
    ...
    route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    route.Run()
}
```