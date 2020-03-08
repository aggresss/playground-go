## dep-demo

下载安装：

```
go get -u -v github.com/golang/dep/cmd/dep
```

特征文件：`Gopkg.lock` `Gopkg.toml`

操作：

```
dep ensure
```

说明：

```
Dep is a tool for managing dependencies for Go projects

Usage: "dep [command]"

Commands:
  init     Set up a new Go project, or migrate an existing one
  status   Report the status of the project's dependencies
  ensure   Ensure a dependency is safely vendored in the project
  prune    Pruning is now performed automatically by dep ensure.
  version  Show the dep version information

Examples:
  dep init                               set up a new project
  dep ensure                             install the project's dependencies
  dep ensure -update                     update the locked versions of all dependencies
  dep ensure -add github.com/pkg/errors  add a dependency to the project

Use "dep help [command]" for more information about a command.

```
