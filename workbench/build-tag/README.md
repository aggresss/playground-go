

https://golang.org/pkg/go/build/#hdr-Build_Constraints

A build constraint is evaluated as the OR of space-separated options. Each option evaluates as the AND of its comma-separated terms.

// +build linux,386 darwin,!cgo

(linux AND 386) OR (darwin AND (NOT cgo))

```
*_GOOS
*_GOARCH
*_GOOS_GOARCH
```

(example: source_windows_amd64.go) where GOOS and GOARCH represent any known operating system and architecture values respectively, then the file is considered to have an implicit build constraint requiring those terms (in addition to any explicit constraints in the file).

## Reference

- [https://golang.org/pkg/go/build/](https://golang.org/pkg/go/build/)
