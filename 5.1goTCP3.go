package main

import (
	"fmt"
	"net"
)

func main() {
	ip, error := net.ResolveTCPAddr("tcp","www.google.com:80")
	if error != nil {
		fmt.Printf("出现错误，错误信息为%v", error)
	}else {
		fmt.Printf("目标TCP地址为：%v", ip)
	}
}
