package main

import (
	"fmt"
	"net"
)

func main() {
	address, error := net.LookupIP("www.google.com")
	if error != nil {
		fmt.Printf("出现错误，错误信息为%v", error)
	}
	fmt.Printf("目标主机地址为：%v", address)
}
