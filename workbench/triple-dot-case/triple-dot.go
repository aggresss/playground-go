// https://programming.guide/go/three-dots-ellipsis.html

package main

import "fmt"

func test1(args ...string) { // variadic function
	for _, v := range args {
		fmt.Println(v)
	}
}

func main() {
	var str1 = []string{
		"qwer1",
		"asdf1",
		"zxcv1",
	}

	var str2 = [...]string{ // 声明一个数组
		"qwer2",
		"asdf2",
		"zxcv2",
	}
	test1(str1...)                   //切片被打散传入
	test1(str2[0], str2[1], str2[2]) // 数组不支持打散传递
}

// 使用 go run ./... 可以遍历当前目录下所有可以运行的 *.go 文件
