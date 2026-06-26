package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeter()
	result :=adder(3,5)
	
	fmt.Println("result is :",result)
	proResult :=proAdder(2,3,6,8,9)
	fmt.Println("pro result is",proResult)
}
func adder(valone int,valtwo int) int {
	return valone+valtwo
}
func proAdder (values ...int)int{
	total := 0
	for _,val:= range values{
		total+=val
	}
	return total
}

 func greeter(){
	fmt.Println("heyy from go")
 }