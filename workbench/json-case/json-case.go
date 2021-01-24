// https://www.jianshu.com/p/31757e530144

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Product for test json
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
	IsRepo    bool    `json:"is_repo"`
	PrivateID int64   `json:"-"` // 表示不进行序列化
	Note      string
}

func test1() {
	var data = `{
		"name":"Xiao mi 6",
		"product_id":"10",
		"number":"10000",
		"price":"2499",
		"is_on_sale":"true",
		"note":"foobar"
		}`

	p1 := &Product{}
	if err1 := json.Unmarshal([]byte(data), p1); err1 != nil {
		fmt.Println(err1)
	}
	fmt.Printf("%#v\n", p1)

	b, err2 := json.Marshal(p1)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Printf("%s\n", b)

	p2 := new(Product)
	if err2 := json.NewDecoder(strings.NewReader(data)).Decode(p2); err2 != nil {
		log.Fatalln(err2)
	}
	fmt.Printf("%#v\n", p2)

}

type A struct {
	X int `json:"X"`
	Z int `json:"Z"`
}

type B struct {
	Y int `json:"Y"`
	Z int `json:"Z"`
}

type User struct {
	ID string
	A
	B
}

func test2() {
	s1 := User{
		ID: "admin",
		A: A{
			X: 1,
			Z: 2,
		},
		B: B{
			Y: 5,
			Z: 6,
		},
	}
	fmt.Printf("%+v\n", s1)

	b, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", b)
	}

	testJson := `{"ID":"admin","X":1,"Y":2,"Z":6}`

	s2 := User{}
	err = json.Unmarshal([]byte(testJson), &s2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", s2)
	}
}

func main() {
	test2()
}
