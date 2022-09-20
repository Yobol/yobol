# Gin 接口路由过滤器使用指南

在其它框架中，负责完成所有或部分接口“通用操作”的特性通常叫做**过滤器**，如在调用接口时判断用户是否登录或具有响应的操作权限。在 `Gin` 中，使用专用名词 `middleware` 来指代。

## [如何使用](https://github.com/gin-gonic/gin#using-middleware)

`gin.New()` 和 `gin.Default()`  都可以创建路由，区别是 `gin.New()` 创建的是没有任何 `middleware` 的路由，而 `gin.Default()` 则会创建带有 `gin.Logger middleware` 和 `gin.Recovery middleware` 的路由。

- `gin.Logger` 会将日志写入到 `gin.DefaultWriter` 即 `os.Stdout`  中，即使设置了环境变量 `GIN_MODE=release`；

- `gin.Recovery` 会从任何 `panic` 中恢复并返回 `500` 响应。

### 使用 `gin.Logger middleware`

`Gin` 支持将日志写入到不同文件中：

```go
func main() {
    // Disable Console Color, you don't need console color when writing the logs to file.
    gin.DisableConsoleColor()

    // Logging to a file.
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f)

    // Use the following code if you need to write the logs to file and console at the same time.
    // gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    router := gin.Default()
    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    router.Run(":8080")
}
```

`Gin` 提供了 `gin.LoggerWithFormatter` 方法帮助我们自定义日志格式：

```go
func main() {
  router := gin.New()

  // LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
  // By default gin.DefaultWriter = os.Stdout
  router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

    // your custom format
    return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
        param.ClientIP,
        param.TimeStamp.Format(time.RFC1123),
        param.Method,
        param.Path,
        param.Request.Proto,
        param.StatusCode,
        param.Latency,
        param.Request.UserAgent(),
        param.ErrorMessage,
    )
  }))

  router.Use(gin.Recovery())

  router.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })

  router.Run(":8080")
}
```

### 使用 `gin.Recovery middleware`

```go
func main() {
  // Creates a router without any middleware by default
  r := gin.New()

  // Global middleware
  // Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
  // By default gin.DefaultWriter = os.Stdout
  r.Use(gin.Logger())

  // Recovery middleware recovers from any panics and writes a 500 if there was one.
  r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
    if err, ok := recovered.(string); ok {
      c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
    }
    c.AbortWithStatus(http.StatusInternalServerError)
  }))

  r.GET("/panic", func(c *gin.Context) {
    // panic with a string -- the custom middleware could save this to a database or report it to the user
    panic("foo")
  })

  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "ohai")
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

### 使用 `gin.BasicAuth middleware`

```go
func Logger() gin.HandlerFunc {
  return func(c *gin.Context) {
    t := time.Now()

    // Set example variable
    c.Set("example", "12345")

    // before request

    c.Next()

    // after request
    latency := time.Since(t)
    log.Print(latency)

    // access the status we are sending
    status := c.Writer.Status()
    log.Println(status)
  }
}

func main() {
  r := gin.New()
  r.Use(Logger())

  r.GET("/test", func(c *gin.Context) {
    example := c.MustGet("example").(string)

    // it would print: "12345"
    log.Println(example)
  })

  // Listen and serve on 0.0.0.0:8080
  r.Run(":8080")
}
```

### 自定义 `middleware`

```go

```