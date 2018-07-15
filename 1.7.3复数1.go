package main

import "fmt"

func main()  {
	var a complex128=complex(145,64)
	var b complex128=complex(12,7)
	fmt.Println(a,b,real(a),real(b),imag(a),imag(b))
	fmt.Println(a*b,real(a*b),imag(a*b))
}
