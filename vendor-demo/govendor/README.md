## govendor-demo

go 依赖包搜索顺序为 `vendor` => `$GOROOT` => `$GOPATH`

下载 govendor

```
go get -u -v github.com/kardianos/govendor
```

特征文件 `vendor/vendor.json`


>Go 1.6以上版本默认开启 `GO15VENDOREXPERIMENT` 环境变量，可忽略该步骤。
通过设置环境变量 `GO15VENDOREXPERIMENT=1` 使用 vendor 文件夹构建文件。
可以选择 `export GO15VENDOREXPERIMENT=1` 或 `GO15VENDOREXPERIMENT=1 go build` 执行编译

Operation:

```shell
govendor init
govendor fetch +missing

```

命令 | 功能
---|---
init | 初始化 vendor 目录
list | 列出所有的依赖包
add | 添加包到 vendor 目录，如 govendor add +external 添加所有外部包
add PKG\_PATH | 添加指定的依赖包到 vendor 目录
update | 从 $GOPATH 更新依赖包到 vendor 目录
remove | 从 vendor 管理中删除依赖
status | 列出所有缺失、过期和修改过的包
fetch | 添加或更新包到本地 vendor 目录
sync | 本地存在 vendor.json 时候拉去依赖包，匹配所记录的版本
get | 类似 go get 目录，拉取依赖包到 vendor 目录

对于 govendor 来说，主要存在三种位置的包：项目自身的包组织为本地（local）包；传统的存放在 $GOPATH 下的依赖包为外部（external）依赖包；被 govendor 管理的放在 vendor 目录下的依赖包则为 vendor 包。
具体来看，这些包可能的类型如下：

状态 | 状态缩写 | 含义
---|:---:|---
+local | l | 本地包，即项目自身的包组织
+external | e | 外部包，即被 $GOPATH 管理，但不在 vendor 目录下
+vendor | v | 已被 govendor 管理，即在 vendor 目录下
+std | s | 标准库中的包
+unused | u | 未使用的包，即包在 vendor 目录下，但项目并没有用到
+missing | m | 代码引用了依赖包，但该包并没有找到
+program | p | 主程序包，意味着可以编译为执行文件
+outside | | 外部包和缺失的包
+all | | 所有的包

参考 `govendor --help`

