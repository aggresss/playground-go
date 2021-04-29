package main

// 由 array 通过 [:] 切片操作后得到的 slice 只有在进行 append 操作后会进行内存复制；
// 由 slice 通过 [:] 切片操作后得到的 slice 在 append 操作后不会进行内存复制；

// slice -> slice 的切片操作不会引起内存复制

// Idioms
// [:] => 创建一个新的引用
// [:0] => reset slice

import "fmt"

func test_01() {
	var bArray [20]int
	// b = append(b, 3) // Array 不能 append

	bSlice := bArray[:]
	bSlice = append(bSlice, 3) // 重新切片后转换为 Slice
}

func test_02() {
	var bArray [20]int
	bSlice := bArray[:]
	bSlice[1] = 3
	bArray[1] = 2

	fmt.Println(bSlice[1]) // 2 如果切片后不 append，则仍然使用不进行内存复制
}

func test_03() {
	var bArray [20]int
	bSlice := bArray[:]
	bSlice = append(bSlice, 0)
	bSlice[1] = 3
	bArray[1] = 2

	fmt.Println(bSlice[1]) // 3 如果切片后有 append 操作，则进行内存复制
}

func test_04() {
	a := make([]int, 0)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	b := a[:]
	b = append(b, 0)
	b[1] = 100

	fmt.Println(a[1], len(a)) // 100 // 仍然引用原始内存
}

func test_05() {
	var a []int = nil
	var a0 []int = make([]int, 0)

	fmt.Println(a == nil)  // true
	fmt.Println(a0 == nil) // false

	fmt.Printf("%#v\n", a)  // []int(nil) // nil slice
	fmt.Printf("%#v\n", a0) // []int{} // empty slice
	// Go wiki recommends using nil slices over empty slices
	// Reference: https://yourbasic.org/golang/clear-slice
	// Note that there are limited circumstances where a non-nil but zero-length slice is preferred,
	// such as when encoding JSON objects
	// (a nil slice encodes to null, while []string{} encodes to the JSON array []).
}

func main() {
	test_04()
}
