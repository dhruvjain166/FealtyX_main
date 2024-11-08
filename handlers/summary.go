package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "unique-student-api/models"
)

// GetStudentSummary retrieves a summary of a student by ID.
func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    if student, ok := models.StudentStore[id]; ok {
        // Generate a summary of the student here (e.g., using Ollama API)
        summary := "Summary of student " + student.Name // Placeholder summary
        json.NewEncoder(w).Encode(map[string]string{"summary": summary})
    } else {
        http.Error(w, "Student not found", http.StatusNotFound)
    }
}
