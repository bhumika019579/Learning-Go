package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
performGetRequest()
performPostJsonRequest()
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
func performPostJsonRequest(){
	const myurl="http://localhost:8003/post"
	//fake json payload
	requestBody:= strings.NewReader(`
	{
	"coursename":"golang",
	"price":0,
	"platform":"lco"
	}
	`)
	response,err:=http.Post(myurl,"application/json",requestBody)
	if err!=nil{
		panic(err)
	}
	defer response.Body.Close()
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))

}