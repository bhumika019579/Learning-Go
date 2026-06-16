package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")
	//var ptr *int
	//fmt.Println("value of pointer is ",ptr)
	mynumber :=23
	var ptr = &mynumber
	fmt.Println("value of pointer is",ptr)
	fmt.Println("value of pointer is",*ptr)
	*ptr=*ptr+2
	fmt.Println("new value is ",mynumber)


}