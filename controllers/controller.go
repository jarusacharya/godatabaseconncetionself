package controller

import (
	handler "databaseserverintegration/handlers"
	"databaseserverintegration/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// get all course
func GetAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appliction/json")
	courses := handler.GetallCourse()
	err := json.NewEncoder(w).Encode(courses)
	if err != nil {
		log.Fatal(err)
	}
}

// get a course
func GetACourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	course := handler.GetoneCourse(params["id"])
	err := json.NewEncoder(w).Encode(course)
	if err != nil {
		log.Fatal(err)
	}

}

// Delete all courses
func DeleteAllCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(handler.DeleteallCourse())
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteaCourse
func DeleteACourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if err := json.NewEncoder(w).Encode(handler.DeleteoneCourse(params["id"])); err != nil {
		log.Fatal(err)
	}
}

// Insert a course
func InsertACourse(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		log.Fatal(err)
	}

	err := json.NewEncoder(w).Encode(handler.InsertoneCourse(course))
	if err != nil {
		log.Fatal(err)
	}

}

//update a course

func UpdateACourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if err := json.NewEncoder(w).Encode(handler.UpdateoneCourse(params["id"])); err != nil {
		log.Fatal(err)
	}

}
