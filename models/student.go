package models

// Student struct holds individual student information.
type Student struct {
    ID    int    `json:"id" validate:"required"`
    Name  string `json:"name" validate:"required,min=2"`
    Age   int    `json:"age" validate:"gte=1,lte=120"`
    Email string `json:"email" validate:"email"`
}

// StudentStore holds the map of students by ID.
var StudentStore = make(map[int]Student)

func init() {
    // Corrected: using StudentStore instead of Students
    StudentStore[1] = Student{ID: 1, Name: "Rohit Sharma", Age: 20, Email: "rohit@example.com"}
    StudentStore[2] = Student{ID: 2, Name: "Virat Kohli", Age: 22, Email: "virat@example.com"}
    StudentStore[3] = Student{ID: 3, Name: "Prithviraj Awatade", Age: 19, Email: "prithviraj@example.com"}
}
