# Gin 请求响应处理

## 请求

### 处理请求参数 `param`

#### 路径参数 `path param` 解析

> [parameters in path](https://github.com/gin-gonic/gin#parameters-in-path)

`Gin` 框架支持路径参数，提供了 `:` 和 `*` 两个特殊符号进行辅助。

- `:`：表示必须包含这部分；

- `*`: 表示可以省略这部分；

`注：确切的路由总会在参数路径之前被匹配。`

```go
func main() {
  router := gin.Default()

  // This handler will match /users/john but will not match /users/ or /users
  router.GET("/users/:name", func(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  // However, this one will match /users/john/ and also /users/john/send
  // If no other routers match /users/john, it will redirect to /users/john/
  router.GET("/users/:name/*action", func(c *gin.Context) {
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })

  // For each matched request Context will hold the route definition
  router.POST("/users/:name/*action", func(c *gin.Context) {
    b := c.FullPath() == "/users/:name/*action" // true
    c.String(http.StatusOK, "%t", b)
  })

  // This handler will add a new router for /users/groups.
  // Exact routes are resolved before param routes, regardless of the order they were defined.
  // Routes starting with /user/groups are never interpreted as /users/:name/... routes
  router.GET("/users/groups", func(c *gin.Context) {
    c.String(http.StatusOK, "The available groups are [...]")
  })

  router.Run(":8080")
}
```

#### 查询参数 `query param` 解析

> [query parameters](https://github.com/gin-gonic/gin#querystring-parameters)

```go
func main() {
    router := gin.Default
    // The request responds to a url matching: /products?category=book&key=gin
    router.Get("/products", func(c *gin.Context) {
        category := c.DefaultQuery("category", "all")
        key := c.Query("") // shortcut for c.Reuqest.URL.Query().Get("key")

        ...
    })
    router.Run(":8080")
}
```

#### 请求体 `request body` & 请求表单 `form` 解析

> [model binding and validation](https://github.com/gin-gonic/gin#model-binding-and-validation)

```go
// Binding from JSON
type Login struct {
  User     string `form:"user" json:"user" xml:"user"  binding:"required"`
  Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
    router := gin.Default()

    // Example for binding JSON ({"user": "manu", "password": "123"})
    router.POST("/loginJSON", func(c *gin.Context) {
        var json Login
        if err := c.ShouldBindJSON(&json); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
        }

        if json.User != "manu" || json.Password != "123" {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
        }

        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    })

    // Example for binding XML (
    //  <?xml version="1.0" encoding="UTF-8"?>
    //  <root>
    //    <user>manu</user>
    //    <password>123</password>
    //  </root>)
    router.POST("/loginXML", func(c *gin.Context) {
        var xml Login
        if err := c.ShouldBindXML(&xml); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
        }

        if xml.User != "manu" || xml.Password != "123" {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
        }

        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    })

    // Example for binding a HTML form (user=manu&password=123)
    router.POST("/loginForm", func(c *gin.Context) {
        var form Login
        // This will infer what binder to use depending on the content-type header.
        if err := c.ShouldBind(&form); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
        }

        if form.User != "manu" || form.Password != "123" {
        c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        return
        }

        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    })

    // Listen and serve on 0.0.0.0:8080
    router.Run(":8080")
}
```

请求示例：

```shell
$ curl -v -X POST \
  http://localhost:8080/loginJSON \
  -H 'content-type: application/json' \
  -d '{ "user": "manu" }'
> POST /loginJSON HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.51.0
> Accept: */*
> content-type: application/json
> Content-Length: 18
>
* upload completely sent off: 18 out of 18 bytes
< HTTP/1.1 400 Bad Request
< Content-Type: application/json; charset=utf-8
< Date: Fri, 04 Aug 2017 03:51:31 GMT
< Content-Length: 100
<
{"error":"Key: 'Login.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
```

## 响应