package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"
	fmt.Println("list of all languages:",languages)
	fmt.Println("JS shorts for :",languages["JS"])
	delete(languages,"RB")
	fmt.Println("list for all languages:",languages)
	//loops are interesting in golang
	for key,value := range languages{
		fmt.Printf("for key %v,value is %v\n",key,value)
	}

}