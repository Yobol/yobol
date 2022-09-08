# Viper 核心概念与常用操作

## 核心概念

[Viper](https://github.com/spf13/viper) 官方将其定义为 `Go configuration with fangs`，可以引申为 Go 配置文件大杀器。具有如下特性：

- 读取本地 `JSON`、`TOML`、`YAML`、`HCL`、`.env` 文件并监听更新事件
  
- 支持读取远程配置文件（etcd 或 Consul）并监听更新事件

`Viper` 可以作为应用的全局配置代理使用。


`Viper` 支持多种方式进行配置，遵循如下优先级（从高到低）：

- 显式调用 `viper.Set()` 方法；

- 命令行参数 `flag.Parse()`；

- 系统环境变量；

- `AddConfigPath` & `SetConfigName` & `SetConfigType` 设置的配置文件；

- 键值存储（etcd/Consul)；

- 默认值 `viper.SetDefault()`；

有一点需要注意的是，`Viper` 配置的键的大小写敏感的。

## 基本操作

### 安装

```shell
go get github.com/spf13/viper
```

### 设置默认值

```Go
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"key1": "value1", "key2": "value2"})
```

### 读取配置文件

```Go
viper.SetConfigName("config") // name of config file (without extension)
viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
viper.AddConfigPath("/etc/appname/")   // path to look for the config file in
viper.AddConfigPath("$HOME/.appname")  // call multiple times to add many search paths
viper.AddConfigPath(".")               // optionally look for config in the working directory
err := viper.ReadInConfig() // Find and read the config file
if err != nil { // Handle errors reading the config file
	panic(fmt.Errorf("fatal error config file: %w", err))
}
```

`Viper` 支持使用 `AddConfigPath` 从多个文件路径寻找配置文件，但是对于每个 `Viper` 实例来说只能由一个配置文件初始化，并且要保证在指定的路径列表下至少有一个配置文件存在。

## 高级特性
