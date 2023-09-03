
## Reference

- https://golang.org/cmd/cgo/
- http://golang.org/doc/articles/c_go_cgo.html
- https://github.com/golang/go/wiki/cgo

- [Go 与 C 的桥梁：cgo 入门，剖析与实践](https://zhuanlan.zhihu.com/p/349197066)
- [Statically Linking Go in 2022](https://mt165.co.uk/blog/static-link-go/)

通过 import “C” 语句启用 CGO 特性后，CGO 会将上一行代码所处注释块的内容视为 C 代码块，被称为序文（preamble）。
