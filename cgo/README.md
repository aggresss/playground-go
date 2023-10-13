
## Reference

- https://golang.org/cmd/cgo/
- http://golang.org/doc/articles/c_go_cgo.html
- https://github.com/golang/go/wiki/cgo

- [Go 与 C 的桥梁：cgo 入门，剖析与实践](https://zhuanlan.zhihu.com/p/349197066)
- [Statically Linking Go in 2022](https://mt165.co.uk/blog/static-link-go/)
- [Hidden Dragons of CGO: Hard-Learned Lessons from Writing Go Wrappers](https://docs.yottadb.com/Presentations/DragonsofCGO.pdf)
-

通过 import “C” 语句启用 CGO 特性后，CGO 会将上一行代码所处注释块的内容视为 C 代码块，被称为序文（preamble）。

C 访问 Go 函数
- c 中 declare
- go 中 definition 并 //export <func_name>
- 使用 (C.type)(C.<func_name>)
