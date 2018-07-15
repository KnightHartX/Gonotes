package main

import "fmt"

func main()  {
	fmt.Println("请输入一个正整数n")
	var n int
	fmt.Scanln(&n)
	m:=n%2
	switch m {
	case 0:fmt.Printf("%v为偶数\n",n)
	case 1:fmt.Printf("%v为奇数\n",n)
	}
}