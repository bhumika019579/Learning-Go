package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("value of a dice is", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("dice value is 1 and you can open ")
	case 2:
		fmt.Println("you can move 2 spots")
	case 3:
		fmt.Println("you can move 3 spots")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spots")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spots")
	case 6:
		fmt.Println("you can move 6 spots and roll dice again")
	default:
		fmt.Println("what was that")
	}
}
