package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
performGetRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8003/get"
	response, err := http.Get(myurl)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code :", response.StatusCode)
	fmt.Println("content length is :",response.ContentLength)
	var responseString strings.Builder
	content,_:=ioutil.ReadAll(response.Body)
	bytecount,_:=responseString.Write(content)
	fmt.Println(bytecount)
	fmt.Println(responseString.String())
}