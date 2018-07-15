package main

import "fmt"

func main()  {
	a:=make([]int,5,5)
	a=[]int{1,2,3,4,5}
	b:=make([]int,5,5)
	copy(b,a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	fmt.Println(b)
	a=append(a,6)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)
	a1:=a[0:6]
	fmt.Println(a1)
}