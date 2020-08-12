## testify

**Reference**
1. [golang的测试框架stretchr/testify](https://www.jianshu.com/p/ad46bbbf877c)
2. [https://github.com/stretchr/testify](https://github.com/stretchr/testify)

```
go get github.com/stretchr/testify
```

```
github.com/stretchr/testify/assert
github.com/stretchr/testify/require
github.com/stretchr/testify/mock
github.com/stretchr/testify/suite
```

断言中 assert 和 require 区别
他们的唯一差别就是 require 的函数会直接导致 case 结束，而 assert 虽然也标记为 case 失败，但 case 不会退出，而是继续往下执行。
