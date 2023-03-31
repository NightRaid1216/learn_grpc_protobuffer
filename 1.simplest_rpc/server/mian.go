package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func add(writer http.ResponseWriter, r *http.Request) {
	//解析参数
	err := r.ParseForm()
	if err != nil {
		return
	}

	/*
		对于POST或PUT请求，ParseForm会将body当作表单解析，并将结果既更新到r.PostForm也更新到r.Form
		ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段。
		ParseMultipartForm会自动调用ParseForm。重复调用本方法是无意义的。

	*/
	//"localhost:8080/add?a=1&b=2"
	fmt.Println("path:", r.URL.Path)

	/*
		a, err := strconv.Atoi(r.Form["a"][0]) 猜想：Form["a"][0] [0]代表空的[]String
	*/

	a, err := strconv.Atoi(r.Form["a"][0])
	if err != nil {
		return
	}
	b, err := strconv.Atoi(r.Form["b"][0])
	if err != nil {
		return
	}

	/*
		Set会添加“Content-Type”："application/json"到Writer
	*/

	writer.Header().Set("Content-Type", "application/json")

	/*
		返回json  "data": a + b,
	*/

	mJson, err := json.Marshal(map[string]int{
		"data": a + b,
	})
	if err != nil {
		return
	}

	/*
		将 mJson 写入缓冲区
	*/

	writer.Write(mJson)

}

func main() {
	http.HandleFunc("/add", add)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
