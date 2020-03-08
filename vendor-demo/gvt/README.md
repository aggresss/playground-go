## gvt

下载安装：

```
go get -u -v github.com/FiloSottile/gvt
```

特征文件：`vendor/manifest`

操作：

```
gvt fetch github.com/sirupsen/logrus

```

说明：

```
gvt, a simple go vendoring tool based on gb-vendor.

Usage:
        gvt command [arguments]

The commands are:

        fetch       fetch a remote dependency
        restore     restore dependencies from manifest
        update      update a local dependency
        list        list dependencies one per line
        delete      delete a local dependency

Use "gvt help [command]" for more information about a command.
```

