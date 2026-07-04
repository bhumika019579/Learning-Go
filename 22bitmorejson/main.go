package main

import (
	"encoding/json"
	"fmt"
)
type course struct{
	Name string `json:"coursename"`
	Price int
	Platform string `json:"website"`
	Password string  `json:"-"` //placing - in any struct prevents it from printing in json as password should not be leaked 
	Tags []string     `json:"tags,omitempty"` //If Tags is empty or nil, don't include it in the JSON
}

func main() {
EncodeJson()
}
func EncodeJson(){
	lcocourses:=[]course{
		{"ReactJS Bootcamp",299,"lco.in","abc123",[]string{"web-dev","js"}},
		{"NextJS Bootcamp",399,"lco.in","bcd123",[]string{"frontend","js"}},
		{"Angular Bootcamp",299,"lco.in","hit123",nil},
	}
	//package this data as json data
	finalJson,err:=json.MarshalIndent(lcocourses,"","\t")
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)

}

