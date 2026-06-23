package main

import "fmt"

func main() {
	// no inheritance in golang ; no super or parent
	bhumika := User{"bhumika", "bhumika019576@gmail.com", true, 21}
	fmt.Println(bhumika)
	fmt.Printf("bhumika details are : %+v\n", bhumika)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
