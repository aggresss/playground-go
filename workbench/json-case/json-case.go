package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func main() {
	var data = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
	p := &Product{}
	err := json.Unmarshal([]byte(data), p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*p)
}
