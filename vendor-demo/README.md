## Golang 常见包管理器

包管理器 | 特征文件 | 说明
:---:|---|---
dep | `Gopkg.lock` `Gopkg.toml` | https://golang.github.io/dep/
govendor | `vendor/vendor.json` | https://github.com/kardianos/govendor
glide | `glide.yaml` `glide.lock` | https://glide.sh/
gvt | `vendor/manifest` | https://github.com/FiloSottile/gvt
go mod | `go.mod` `go.sum` | https://blog.golang.org/migrating-to-go-modules

>Go 1.6以上版本默认开启 `GO15VENDOREXPERIMENT` 环境变量，可忽略该步骤。
通过设置环境变量 `GO15VENDOREXPERIMENT=1` 使用 vendor 文件夹构建文件。
可以选择 `export GO15VENDOREXPERIMENT=1` 或 `GO15VENDOREXPERIMENT=1 go run` 执行编译

