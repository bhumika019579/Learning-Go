package main

import (

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
   "math/rand"

	"github.com/gorilla/mux"
)

// model for courses-file
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`

}
type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}
//fake db
var courses []Course
// middleware,helper-file
func(c*Course) IsEmpty() bool{
//return c.CourseId=="" && c.CourseName==""
return c.CourseName==""
}
func main(){

}
// controllers-file 
//serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>welcome to API by learncode online</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all courses ")
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Get one course")
	w.Header().Set("content-type","application/json")
	//grab id from request
	params:=mux.Vars(r)
    //loop through courses,find matching id and return the response
	for _,Course:=range courses{
		if Course.CourseId==params["id"]{
			json.NewEncoder(w).Encode(Course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with given id ")
	return 

}
func createOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("create one course")
	w.Header().Set("content-type","application/json")
	// what if : body is empty
	if r.Body==nil{
		json.NewEncoder(w).Encode("please send some data")

	}
	// what about-{}
	var course Course
	_=json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("please send some data")
		return 
	}
	// generate unique id ,string 
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId= strconv.Itoa(rand.Intn(100))
	courses=append(courses,course)
	json.NewEncoder(w).Encode(course)
	return 

}