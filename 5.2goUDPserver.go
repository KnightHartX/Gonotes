package main

import
(
"fmt"
"net"
)

func HandleCilent(conn *net.UDPConn,data []byte,address *net.UDPAddr)  {
	fmt.Println("收到数据："+string(data))
	conn.WriteToUDP([]byte("数据已收到"),address)
}

func main() {
	//解析UDP地址
	address, error := net.ResolveUDPAddr("udp", ":7070")
	if error != nil {
		fmt.Printf("出现错误，错误信息为%v", error)
	} else {
		return
	}
	//7070端口监听
	conn, error := net.ListenUDP("udp", address)
	if error != nil {
		fmt.Printf("出现错误，错误信息为%v", error)
	} else {
		return
	}
	for{//循环接收数据，处理数据
		var buf [1024]byte
		n,address,error:=conn.ReadFromUDP(buf[0:])
		if error != nil {
			fmt.Printf("出现错误，错误信息为%v", error)
		} else {
			return
			//开启新线程处理客户端数据
			go HandleCilent(conn,buf[0:n],address)
	}
}
}