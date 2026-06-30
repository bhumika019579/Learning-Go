package main

import (
	"fmt"
	"net/http"
)

func main() {
performGetRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code :", response.StatusCode)
	fmt.Println("content length is :",response.ContentLength)
}