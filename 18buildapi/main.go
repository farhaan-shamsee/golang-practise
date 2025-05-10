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

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - usually goes into seperate file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == "" //simple way
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to API building")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS",
		CoursePrice: 299, Author: &Author{Fullname: "Farrow", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack",
		CoursePrice: 299, Author: &Author{Fullname: "FarrowAman", Website: "go.dev"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal((http.ListenAndServe(":4002", r)))

}

// controllers file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welsome to API by golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	fmt.Printf("Type pf params is: %T\n", params)
	fmt.Printf("Content of params is: %v\n", params)

	// loop through course, find matching id and return response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if: data is send like {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	_ = json.NewDecoder(r.Body).Decode(&course)
	// if courseID is already there
	var duplicateEntry = false
	for _, courseVar := range courses {
		if courseVar.CourseName == course.CourseName {
			duplicateEntry = true
			break
		}
	}

	if duplicateEntry {
		json.NewEncoder(w).Encode("Duplicate entry")
	} else {
		// generate unique ID, string
		// append course into courses
		rand.Seed(time.Now().UnixNano())
		course.CourseId = strconv.Itoa(rand.Intn(100))
		courses = append(courses, course)
		json.NewEncoder(w).Encode(courses)

	}

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	var courseFound = false
	// loop, id, remove, add with myId
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			courseFound = true
			json.NewEncoder(w).Encode("Course Update")
			break
		}
	}

	// TODO: send a response when id is not found
	// for _, course := range courses {
	// 	if course.CourseId != params["id"] {
	// 		json.NewEncoder(w).Encode("No course found with mentioned ID: " + params["id"])
	// 	}
	// }
	if !courseFound {
		json.NewEncoder(w).Encode("No course found with mentioned ID: " + params["id"])
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("The course is deleted with ID:" + params["id"])
			break
		}
	}

}



// Example Calls
// curl --request POST "http://localhost:4002/course" \
//      --header "Content-Type: application/json" \
//      --data '{
//        "courseid": "101",
//        "coursename": "Introduction to Go",
//        "price": 99,
//        "author": {
//          "fullname": "John Doe",
//          "website": "https://johndoe.com"
//        }
//      }'