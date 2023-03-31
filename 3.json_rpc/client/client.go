package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
1.建立连接
2.调用函数
*/

/*
	{
	"method":"HS.Hello",
	"params":"bobby",
	"id":0
	}

*/

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	//var reply *string = new(string) //如果不赋值，是空指针，无法给reply分配数值
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HS.Hello", "yuTong", &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}
