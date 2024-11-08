package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "unique-student-api/models"
)

// UpdateStudent updates the details of a student by ID.
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var updatedStudent models.Student
    if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
        http.Error(w, "Invalid data", http.StatusBadRequest)
        return
    }

    if _, ok := models.StudentStore[id]; ok {
        updatedStudent.ID = id
        models.StudentStore[id] = updatedStudent
        json.NewEncoder(w).Encode(updatedStudent)
    } else {
        http.Error(w, "Student not found", http.StatusNotFound)
    }
}
