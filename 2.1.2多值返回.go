package main

import "fmt"

func sqare(a float64) (S float64,ok bool) {
	const pi  =3.1415926
	S=pi*a*a
	ok=true
	return S,ok
}

func main()  {
	var r float64
	fmt.Println("请输入半径r")
	fmt.Scan(&r)
	fmt.Println(sqare(r))
}