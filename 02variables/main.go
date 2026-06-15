package main

import "fmt"

func main() {
	var username string ="bhumika"
	fmt.Println(username)
	fmt.Printf("variable is of type:%T \n",username)
	var smallval uint8=255
	fmt.Println(smallval)
var smallfloat float32 =255.678945678
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type:%T \n",smallfloat)
	//default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type : %T\n",anothervariable) //gives 0
	//no var style 
	numberofusers :=3000
	fmt.Println(numberofusers) //not globally allowed
}
