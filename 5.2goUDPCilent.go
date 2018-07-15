package main

import (
	"net"
	"fmt"
)

func main()  {
	//解析服务器UDP地址
	address, error := net.ResolveUDPAddr("udp", "127.0.0.1:7070")
	if error != nil {
		fmt.Printf("出现错误，错误信息为%v", error)
	} else {
		return
	}
	//连接服务器
	conn,error:=net.DialUDP("udp",nil,address)
	if error != nil {
		fmt.Printf("出现错误，错误信息为%v", error)
	} else {
		return
	}
	defer conn.Close()//关闭连接
	//向服务器发送数据
	conn.Write([]byte("Hello,Server"))
	var buf [1024]byte
	//读取服务器相应信息
	n,_,error:=conn.ReadFromUDP(buf[0:])
	if error != nil {
		fmt.Printf("出现错误，错误信息为%v", error)
	} else {
		return
	}
	fmt.Println(string(buf[0:n]))
}