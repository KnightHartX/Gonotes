package main

import "fmt"

func main()  {
	a:=make(map[string]int)
	a["china"]=1
	a["us"]=2
	fmt.Println(a["us"])
	delete(a,"us")
	a["jp"]=3
	for k, v := range a {
		fmt.Printf("%v is no.%v\n", k, v)
	}
}