package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "unique-student-api/models"
)

// GetStudent retrieves a student by ID.
func GetStudent(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if student, ok := models.StudentStore[id]; ok {
        json.NewEncoder(w).Encode(student)
    } else {
        http.Error(w, "Student not found", http.StatusNotFound)
    }
}
