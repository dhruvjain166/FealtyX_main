package handlers

import (
    "net/http"
    "strconv"
    "unique-student-api/models"
)

// RemoveStudent deletes a student by ID.
func RemoveStudent(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    // Check if the student exists in the StudentStore
    if _, ok := models.StudentStore[id]; ok {
        // Delete the student from the StudentStore
        delete(models.StudentStore, id)
        w.WriteHeader(http.StatusNoContent)
        return
    }

    // If student not found, return a 404 error
    http.Error(w, "Student not found", http.StatusNotFound)
}
