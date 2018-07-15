package main

import "fmt"

type circle struct
{
 	r float64
	h float64
}

type mass struct {
	circle
	D float64
}

func S(p circle) float64 {
	pi:=3.1415926
	return pi*p.r*p.r*p.h
}

func M(p mass) float64 {
	pi:=3.1415926
	return p.h*p.r*p.r*pi*p.D
}

func SS(p *circle) float64 {
	pi:=3.1415926
	return pi*p.r*p.r*p.h
}

func main()  {
	circle1:=circle{2.4,3.1}
	fmt.Println(circle1)
	mass1:=mass{circle{2.4 ,3.1},3.23}
	fmt.Println(mass1)
	fmt.Println(S(circle1))
	fmt.Println(M(mass1))
	fmt.Println(SS(&circle1))
}