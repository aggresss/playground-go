## go-mod-demo

### Operation

```shell
go mod init
go mod vendor
```
相对于 `govendor`，`go mod` 会自动分析当前模块的外部依赖并缓存到 `$GOPATH/pkg/` 下面，省去了手动  `go get` 的过程


### GO111MODULE

Go 1.11 开始支持此功能，建议使用 Go 1.12 以及以上版本

GO111MODULE 有三个值：off, on和auto（默认值）。

- GO111MODULE=off，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。
- GO111MODULE=on，go命令行会使用modules，而一点也不会去GOPATH目录下查找。
- GO111MODULE=auto，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：
	- 当前目录在GOPATH/src之外且该目录包含go.mod文件
	- 当前文件在包含go.mod文件的目录下面。

当modules 功能启用时，依赖包的存放位置变更为`$GOPATH/pkg/mod/`，允许同一个 package 多个版本并存，且多个项目可以共享缓存的 module。


### go mod

golang 提供了 go mod命令来管理包。go mod 有以下命令：

命令 | 说明
---|---
download | download modules to local cache(下载依赖包)
edit | edit go.mod from tools or scripts（编辑go.mod)
graph | print module requirement graph (打印模块依赖图)
init | initialize new module in current directory（在当前目录初始化mod）
tidy | add missing and remove unused modules(拉取缺少的模块，移除不用的模块)
vendor | make vendored copy of dependencies(将依赖复制到vendor下)
verify | verify dependencies have expected content (验证依赖是否正确）
why | explain why packages or modules are needed(解释为什么需要依赖)

参考 `go help mod`
