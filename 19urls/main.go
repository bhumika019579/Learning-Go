package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://ai-audit-seven-snowy.vercel.app/audit/0fh-R-r_"

func main() {
	fmt.Println(myurl)
	//parsing 
	result,_:=url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	partsofurl:=&url.URL{
		Scheme: "https",
		Host: "chatgpt.com",
		Path: "/c/6a425e4b-7e68-83ee-8432-46d2acc9c8b8",
	}
	anotherURL:=partsofurl.String()
	fmt.Println(anotherURL)
}