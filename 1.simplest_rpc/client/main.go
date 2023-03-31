package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func add(a, b int) int {
	req := HttpRequest.NewRequest()
	get, err := req.Get(fmt.Sprintf("http://localhost:8080/%s?a=%d&b=%d", "add", a, b))
	if err != nil {
		panic(err)
	}
	body, err := get.Body()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	rspData := ResponseData{}
	//json 解码成ResponseData
	err = json.Unmarshal(body, &rspData)
	if err != nil {
		panic(err)
	}
	return rspData.Data

}

//RPC ,就像本地函数一样调用服务器上的函数
func main() {
	add(56, 100)
}
