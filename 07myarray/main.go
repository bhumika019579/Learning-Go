package main

import "fmt"

func main() {
	fmt.Println("welcome to arrays")
	var fruitlist [4]string
	fruitlist[0]="apple"
	fruitlist[1]="mango"
	fruitlist[2]="litchi"
	fruitlist[3]="banana"
	fmt.Println("fruit list is :",fruitlist)
	var veglist= [3]string("potato","beans","mushroom")
	fmt.Println("vegy list is:",veglist)
}