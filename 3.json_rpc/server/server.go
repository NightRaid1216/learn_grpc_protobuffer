package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*

 */

type HelloService struct {
}

//不写这个函数会报错,作用是
func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello" + request
	return nil
}

func main() {
	/*
		1.实例化一个 server
		2.注册一个 handler
		3.启动服务
	*/

	//1.实例化一个 server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	//2.RPC 注册一个 handler

	/*
		RegisterName类似Register，但使用提供的name代替rcvr的具体类型名作为服务名。
	*/
	err = rpc.RegisterName("HS", &HelloService{})
	if err != nil {
		panic(err)
	}

	//3.启动服务

	/*
		Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。
	*/

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		//唯一的主要差别 避免多个请求过来 可以开个协程
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
