package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{"apple","tomato","peach"}
	fmt.Printf("type of fruitlist is %T\n",fruitlist)
	fruitlist=append(fruitlist,"mango","banana")
	fmt.Println(fruitlist)
	fruitlist=append(fruitlist[1:3])
	fmt.Println(fruitlist)
	highscores :=make([]int,4)
	highscores[0]=234
	highscores[1]=945
	highscores[2]=465
	highscores[3]=867
	highscores = append(highscores, 555,666,789)
	fmt.Println(highscores)
	sort.Ints(highscores)
	fmt.Println(highscores)
	//how to remove a value from slices based on index
	var courses=[]string{"react.js","javascript","swift","python","ruby"}
	fmt.Println(courses)
	var index int =2
	courses =append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}