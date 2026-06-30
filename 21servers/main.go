package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go")
	})

	fmt.Println("Server started at port 8003")

	err := http.ListenAndServe(":8003", nil)
	if err != nil {
		fmt.Println(err)
	}
}