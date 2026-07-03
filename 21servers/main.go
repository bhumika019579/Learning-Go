package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go")
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		fmt.Println(string(body)) // Prints received data in terminal
		fmt.Fprintln(w, string(body))
	})
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form) //acts like console.log
		fmt.Fprintln(w, r.Form) //acts like res.send
	})

	fmt.Println("Server started at port 8003")

	err := http.ListenAndServe(":8003", nil)
	if err != nil {
		fmt.Println(err)
	}
}
