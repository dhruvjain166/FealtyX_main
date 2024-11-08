package routes

import (
    "github.com/gorilla/mux"
    "unique-student-api/handlers"
)

// SetupRoutes initializes all the routes for the API.
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Route for creating a new student
    router.HandleFunc("/students", handlers.AddStudent).Methods("POST")
    
    // Route for fetching all students
    router.HandleFunc("/students", handlers.GetAllStudents).Methods("GET")
    
    // Route for fetching a single student by ID
    router.HandleFunc("/students/{id}", handlers.GetStudent).Methods("GET")
    
    // Route for updating a student by ID
    router.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")
    
    // Route for deleting a student by ID
    router.HandleFunc("/students/{id}", handlers.RemoveStudent).Methods("DELETE")
    
    // Route for generating a student summary by ID
    router.HandleFunc("/students/{id}/summary", handlers.GetStudentSummary).Methods("GET")

    return router
}
