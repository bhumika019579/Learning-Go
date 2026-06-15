package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to userinput"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza :\n")
// what this reader reads store it into a variable using comma ok syntax||err ok
input, _ :=reader.ReadString('\n')
fmt.Println("thanks for rating, ",input)
fmt.Printf("type of this rating is %T\n ",input)
}