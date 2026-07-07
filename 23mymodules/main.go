package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r:=mux.NewRouter()
	r.HandleFunc("/",serveHome).Methods("Get")
	log.Fatal(http.ListenAndServe(":4001",r))
}
func greeter(){
	fmt.Println("Hey there mod users")
}
func serveHome(w http.ResponseWriter, r*http.Request){
	w.Write([]byte("<h1> welcome to golang </h1>"))
}