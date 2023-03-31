package main

import (
	"fmt"
	"net/rpc"
)

/*
1.建立连接
2.调用函数
*/

func main() {

	dial, err := rpc.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	//var reply *string = new(string) //如果不赋值，是空指针，无法给reply分配数值
	var reply string
	err = dial.Call("HS.Hello", "yuTong", &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}
