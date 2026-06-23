package main

import "fmt"

func main() {
	days := []string{"sunday", "tuesday", "wednesday", "friday", "saturday"}
	fmt.Println(days)
	//for d := 0; d < len(days); d++ {
		//fmt.Println(days[d])

	//}
	for i:=range days{
		fmt.Println(days[i])
	}
	for index,day:=range days{
		fmt.Printf("index is %v and value is %v",index,day) // for each loop in golang
	}
	rouguevalue:=1
	for rouguevalue<10 {
		if rouguevalue==5{
			rouguevalue++
			continue
		}
		fmt.Println("value is :",rouguevalue)
		rouguevalue++
	}
	
}
