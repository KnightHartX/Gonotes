package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main()  {
	var x,y int
	fmt.Println("请输入两个整数")
	fmt.Scanln(&x,&y)
	fmt.Printf("%v和%v的最大公约数为%v",x,y,gcd(x,y))
}
