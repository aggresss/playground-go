// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.6.md

// 通过下面的代码，你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，
// Go通过interface实现了duck-typing:即"当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子"。

package main

import "fmt"

// Men interface
type Men interface {
	SayHi()
	Sing(lyrics string)
}

// Human struct
type Human struct {
	name  string
	age   int
	phone string
}

//SayHi Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Sing Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Student struct
type Student struct {
	Human  //匿名字段
	school string
	loan   float32
}

// Employee struct
type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

// SayHi Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Lawyer struct
// 我们把interface作为struct的一个匿名成员
// 就可以假设struct就是此成员interface的一个实现，而不管struct是否已经实现interface所定义的函数
type Lawyer struct {
	office string
	level  int
	Men
}

// SayHi lawyer say
func (l Lawyer) SayHi() {
	fmt.Printf("Hi, I an a lawyer, I work at %s, my level is: %d\n", l.office, l.level)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
	alex := Lawyer{"TooExpensiveOffice", 3, nil}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	// x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	// x[0], x[1], x[2] = paul, sam, mike
	x := [...]Men{paul, sam, mike, tom, alex}

	for _, value := range x {
		value.SayHi()
	}
}
