# iOS-Tagent Helper

此工具帮助更容易的管理运行 iOS-Tagent。

通常你需要在两个终端窗口中分别运行 xcodebuild 和 iproxy，
现在使用iOS-Tagent Helper，可以使用 tagent 一次启停 xcodebuild 和 iproxy。

前置条件：
- 了解 golang。
- 已通过手工成功配置和运行过 iOS-Tagent。

使用说明：

1. 在 `~/.tagent.yaml` 中配置 iOS-Tagent 的路径

示例：

```yaml
tagent: /Users/mac/iOS-Tagent/WebDriverAgent.xcodeproj
```

2. 编译为可执行文件 tagent

```shell script
$ cd path/to/iOS-Tagent-Helper/tagent
$ go build -o tagent
```

3. 使用 tagent

```shell script
$ ./tagent help # 查看帮助
$ ./tagent devices  # 输出当前 Mac 连接的 iPhone
$ ./tagent connect  # 运行 iOS-Tagent 和 iproxy
$ ./tagent disconnect # 停止运行 iOS-Tagent 和 iproxy
```

其他：

- `--config` 选项可以指定不同路径的配置文件。
- 需要全局可运行，可以复制可执行文件到 `/usr/local/bin` 下，`sudo mv ./tagent /usr/local/bin/`