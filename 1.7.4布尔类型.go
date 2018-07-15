package main

import "fmt"

func main()  {
	var a bool=true
	var b bool=false
	c:=a||b
	d:=a&&b
	fmt.Println(a,b,c,d,convert(c),convert(d))
}

func convert(a bool) int {
	var i int
	if a{
		i=1
	}else {
		i=0
	}
	return i
}