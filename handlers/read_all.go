package handlers

import (
    "encoding/json"
    "net/http"
    "unique-student-api/models"
)

// FetchAllStudents retrieves all students in a formatted JSON response.
func FetchAllStudents(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var students []models.Student
    for _, student := range models.StudentStore {
        students = append(students, student)
    }

    // Format the JSON output with indentation
    jsonResponse, err := json.MarshalIndent(students, "", "  ")
    if err != nil {
        http.Error(w, "Error formatting JSON", http.StatusInternalServerError)
        return
    }

    w.Write(jsonResponse)
}
