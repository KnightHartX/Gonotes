package main

import "fmt"

func main()  {
	var s string="helloworld"
	fmt.Println(s)
	lenth:=len(s)
	fmt.Println(lenth)
	fmt.Println(string(s[1]))
	for i:=0;i<len(s);i++ {
		fmt.Println(string(s[i]))
	}
	s1:=s[1:7]
	s2:=s[0:6]
	println(s1+s2)
}