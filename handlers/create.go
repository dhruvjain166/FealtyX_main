package handlers

import (
    "encoding/json"
    "net/http"
    "sync/atomic"
    "unique-student-api/models"
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()
var currentID int32 = 1

// AddStudent processes POST requests to add a new student.
func AddStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student

    // Decode the request body into the student struct
    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    // Validate the struct
    if err := validate.Struct(student); err != nil {
        http.Error(w, "Validation failed", http.StatusBadRequest)
        return
    }

    // Assign a unique ID to the new student and add it to the StudentStore
    student.ID = int(atomic.AddInt32(&currentID, 1))
    models.StudentStore[student.ID] = student

    // Return a success response
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(student)
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
    // Set the response header to JSON
    w.Header().Set("Content-Type", "application/json")
    
    // Return a list of all students from the StudentStore as JSON
    students := make([]models.Student, 0, len(models.StudentStore))
    
    // Convert the map to a slice of students
    for _, student := range models.StudentStore {
        students = append(students, student)
    }

    // Respond with a 200 OK status and the JSON list of students
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(students)
}
