package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	fmt.Println("API-learncodeonline.in")
	r:=mux.NewRouter()
	// seeding
	courses =append(courses, Course{CourseId: "2",CourseName: "ReactJS",
	CoursePrice: 299,Author:&Author{Fullname: "bhumika",Website: "lco.in"},
})
	courses =append(courses, Course{CourseId: "4",CourseName: "NextJS",
	CoursePrice: 399,Author:&Author{Fullname: "default",Website: "go.in"},
})
//routing
    r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")

// listen to a port 
    log.Fatal(http.ListenAndServe(":5001",r))
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
		return

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
func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("update one course")
	w.Header().Set("content-type","application/json")
	// first -grab id from req
	params:=mux.Vars(r)
	// loop,id ,remove operation,add with my ID 
	for index,SingleCourse:=range courses{
		if SingleCourse.CourseId==params["id"]{
		courses =append(courses[:index],courses[index+1:]...)
		var updatedCourse Course 
		_=json.NewDecoder(r.Body).Decode(&updatedCourse)
		updatedCourse.CourseId=params["id"]
		courses =append(courses, updatedCourse)
		json.NewEncoder(w).Encode("value updated ")
		return
		}
	}
}
func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("delete one course")
	w.Header().Set("content-type","application/json")
	params :=mux.Vars(r)
	//loop ,id,remove(index,index+1)
	for index,OneCourse:=range courses{
		if OneCourse.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]... )
		    json.NewEncoder(w).Encode("deleted successfully")
			return 
           
		}
	}
}