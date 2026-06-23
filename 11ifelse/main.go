package main

import "fmt"

func main() {
	logincount := 23
	var result string
	if logincount < 10 {
		result = "regular user"
	} else if logincount>10 {
		result="watch out"
	} else {
		result="exactly 10 login count"
	}
	fmt.Println(result)
	if 9%2==0{
		fmt.Println("no is even")
	} else {
		fmt.Println("no is odd")
	}
	if num:=3;num<10{
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is more than 10")
	}

}