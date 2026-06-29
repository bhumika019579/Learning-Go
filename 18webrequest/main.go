package main

import (
	"fmt"
	"io"
	"net/http"
)
const url="https://google.com"

func main() {
	fmt.Println("LCO web request")
	response,err:= http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("response is of type : %T\n",response)
	 defer response.Body.Close() //caller 's reponsibility to close the connection
	databytes,err:= io.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(databytes))
	

}