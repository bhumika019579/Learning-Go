package main

import "fmt"

func main() {
	// no inheritance in golang ; no super or parent
	bhumika := User{"bhumika", "bhumika019576@gmail.com", true, 21}
	fmt.Println(bhumika)
	fmt.Printf("bhumika details are : %+v\n", bhumika)
	bhumika.GetStatus()
	bhumika.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
func (u User) GetStatus(){
fmt.Println("Is user active:",u.Status)
}
func (u User) NewMail(){
u.Email="test@go.dev"
fmt.Println("email of this user is :",u.Email)
}