package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"
	fmt.Println("list of all languages:",languages)

}