package main

import "fmt"

func main() {
	fmt.Println("hello")
	defer fmt.Println("world")
	defer fmt.Println("hiee")
	defer fmt.Println("myy")
	fmt.Println("NoOne")
	myDefer()
}
func myDefer(){
	for i:=0;i<5;i++ {
		defer fmt.Println(i)
	}
}