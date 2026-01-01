package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"     // for generating random numbers
	"net/http"      // to build web APIs
	"strconv"       // for converting numbers to strings
	"time"

	"github.com/gorilla/mux" // router package for handling routes
)

// Course is a model or struct that defines how a course looks
type Course struct {
	CourseId    string  `json:"courseid"`   // ID of the course
	CourseName  string  `json:"coursename"` // Name of the course
	CoursePrice int     `json:"price"`      // Price of the course
	Author      *Author `json:"author"`     // Pointer to Author struct
}

// Author struct holds information about course author
type Author struct {
	Fullname string `json:"fullname"` // Author's full name
	Website  string `json:"website"`  // Author's website
}

// This acts like a fake database (in-memory slice of courses)
var courses []Course

// Helper method to check if course data is empty, middleware, helper - usually goes into separate file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == "" //simple way
	// returns true if both ID and Name are empty
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to API building")

	// Create a new router using gorilla/mux
	r := mux.NewRouter()

	// Seed the "database" with some courses
	courses = append(courses, Course{
		CourseId: "2", CourseName: "ReactJS",
		CoursePrice: 299, Author: &Author{Fullname: "Farrow", Website: "lco.dev"},
	})
	courses = append(courses, Course{
		CourseId: "4", CourseName: "MERN Stack",
		CoursePrice: 299, Author: &Author{Fullname: "FarrowAman", Website: "go.dev"},
	})

	// Define API routes and map them to handler functions
	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Start the HTTP server on port 4002
	log.Fatal((http.ListenAndServe(":4002", r)))
}

// --- Controller Functions ---

// serveHome returns a simple HTML page
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by golang</h1>"))
}

// getAllCourses returns all courses in JSON format
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json") // set response type
	json.NewEncoder(w).Encode(courses)                // encode and send all courses
}

// getOneCourse returns a specific course based on ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// Get route parameters (like {id})
	params := mux.Vars(r)

	// Loop through all courses to find matching ID
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course) // send matching course
			return
		}
	}
	// If course not found, return message
	json.NewEncoder(w).Encode("No course found with id")
}

// createOneCourse adds a new course from POST body
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// Check if request body is empty
	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// Decode JSON body into a Course struct
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	
	// Check if course has required fields
	// what if: data is send like {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	// Check for duplicate course name
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
		// Create a new unique Course ID
		rand.Seed(time.Now().UnixNano())
		course.CourseId = strconv.Itoa(rand.Intn(100)) // random number as ID

		// Add course to list
		courses = append(courses, course)

		// Return updated course list
		json.NewEncoder(w).Encode(courses)
	}
}

// updateOneCourse modifies an existing course by ID
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get ID from route

	var courseFound = false
	// Loop through courses to find a match
	for index, course := range courses {
		if course.CourseId == params["id"] {
			// Remove old course
			courses = append(courses[:index], courses[index+1:]...)

			// Add updated course
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			updatedCourse.CourseId = params["id"]
			courses = append(courses, updatedCourse)

			courseFound = true
			json.NewEncoder(w).Encode("Course Updated")
			break
		}
	}

	if !courseFound {
		json.NewEncoder(w).Encode("No course found with mentioned ID: " + params["id"])
	}
}

// deleteOneCourse removes a course by ID
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Loop through courses to find and delete matching course
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