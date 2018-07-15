package main

import "fmt"

func main() {
	S := 0
	i := 0
	n := 100
	for i := 0; i <= n; i++ {
		S = S + i
	}
	println(S)
	i = 0
	for i <= n {
		S = S + i
		i++
	}
	fmt.Println(S)
}
