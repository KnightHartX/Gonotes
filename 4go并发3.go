package main

import "fmt"

func main()  {
	responses := make(chan int, 9)
	go func() {responses<-jisuan(3)}()
	go func() {responses<-jisuan(4)}()
	go func() {responses<-jisuan(5)}()
	go func() {responses<-jisuan(6)}()
	go func() {responses<-jisuan(7)}()
	go func() {responses<-jisuan(8)}()
	go func() {responses<-jisuan(9)}()
	go func() {responses<-jisuan(10)}()
	go func() {responses<-jisuan(11)}()
	a:=<-responses
	fmt.Println(a)
}

func jisuan(a int) int {
	s:=a*a
	return s
}