package main

import "fmt"

func sqare(a float64) float64 {
	const pi  =3.1415926
	S:=pi*a*a
	return S
}

func main()  {
	var r float64
	fmt.Println("请输入半径r")
	fmt.Scan(&r)
	a:=sqare(r)
	b:=sqare(sqare(r))
	fmt.Println(a)
	fmt.Println(b)
}