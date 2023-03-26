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

// Model
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname     string `json:"fullname"`
	Socialhandle string `json:"socialhandle"`
}

// fake db
var courses []Course

// helper function
func (c Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("Building mock api")

	// creating router
	r := mux.NewRouter()

	// seeding fake data
	courses = append(courses, Course{
		CourseId: "22", CourseName: "Android", CoursePrice: 499, Author: &Author{Fullname: "Baklol teacher", Socialhandle: "baklolteacher.com"},
	})
	courses = append(courses, Course{
		CourseId: "33", CourseName: "Next js", CoursePrice: 299, Author: &Author{Fullname: "Black adam", Socialhandle: "blackadam.com"},
	})
	courses = append(courses, Course{
		CourseId: "44", CourseName: "Tailwind", CoursePrice: 99, Author: &Author{Fullname: "Vino", Socialhandle: "vino.com"},
	})

	// decalring route
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourseById).Methods("DELETE")

	// listening port
	log.Fatal(http.ListenAndServe(":8000", r))

}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Busdik</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses baby")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get single course baby")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {

		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create single course baby")
	w.Header().Set("Content-Type", "application/json")

	// if req body is isempty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Bhai data send krne bhul gye ")
		return
	}

	// if data sent is -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Bhai data send kro")
		return
	}

	// generate random number
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update single course baby")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete single course baby")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("deleted")
			break
		}
	}
}
