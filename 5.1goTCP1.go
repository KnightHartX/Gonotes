package main

import (
	"fmt"
	"net"
)

func main()  {
	address,error:=net.InterfaceAddrs()
	if error!=nil{
		fmt.Printf("出现错误，错误信息为%v",error)
	}
	fmt.Printf("本机地址为：%v",address)
}