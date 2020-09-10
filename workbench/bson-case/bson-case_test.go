package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Name  string `json:"name,omitempty" bson:"name,omitempty"`
	CSDN  string `json:"csdn,omitempty" bson:"csdn,omitempty"`
	Quote string `json:"quote,omitempty" bson:"quote,omitempty"`
}

var user = User{
	Name:  "许大侠",
	CSDN:  "https://blog.csdn.net/halo_hsuh",
	Quote: "听我一言, 看清远方",
}

func toBsonBytes() []byte {
	/* 结构体序列化成bson格式[]byte */
	data, err := bson.Marshal(user)
	if nil != err {
		fmt.Println("序列化Bson失败")
		return nil
	}
	return data
}

func toJsonBytes() []byte {
	/* 结构体序列化成bson格式[]byte */
	data, err := json.Marshal(user)
	if nil != err {
		fmt.Println("序列化Json失败")
		return nil
	}

	return data
}

func ToFile(fileName string, data []byte) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if nil != err {
		return err
	}

	length, err := file.Write(data)
	if nil != err {
		return err
	}

	if length < len(data) {
		return errors.New("Write Error: Length Error.")
	}

	return nil
}

func Print(u User) {
	fmt.Println("名字: ", u.Name,
		"; CSDN: ", u.CSDN,
		"; 名言: ", u.Quote)
}

func TestBson(t *testing.T) {
	var tempUser User

	data := toBsonBytes()

	/* bson格式[]byte 反序列化成结构体 */
	err := bson.Unmarshal(data, &tempUser)
	if nil != err {
		fmt.Println("反序列化Bson失败", err)
		return
	}

	Print(tempUser)

	ToFile("./hello.bson", data) // 将序列化的数据存储进文件
}

func TestJson(t *testing.T) {
	var tempUser User

	data := toJsonBytes()

	/* bson格式[]byte 反序列化成结构体 */
	err := json.Unmarshal(data, &tempUser)
	if nil != err {
		fmt.Println("反序列化Json失败: ", err)
		return
	}

	Print(tempUser)

	ToFile("./hello.json", data) // 将序列化的数据存储进文件
}

func TestJsonToBson(t *testing.T) {
	// 获取json []byte
	data := toJsonBytes()

	fmt.Println(data)
	var tempUser User
	// 将json转化成bson 结构体
	err := bson.UnmarshalJSON(data, &tempUser)
	if nil != err {
		fmt.Println("Json 转 Bson失败: ", err)
	}

	Print(tempUser)

	// 结构体再转化为json
	data2, err := bson.MarshalJSON(tempUser)
	if nil != err {
		fmt.Println("Bson 转 Json失败: ", err)
	}

	fmt.Println(string(data2))

}
