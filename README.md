This is a REST API built with Go (Golang) that allows for managing student data. The API supports operations such as adding, updating, deleting, and fetching student details.

Features
Add new student: Create new student records.
Fetch all students: Retrieve a list of all students.
Fetch a student by ID: Retrieve details of a specific student using their unique ID.
Update student details: Update student information such as name, age, and email.
Delete student: Delete a student record by their ID.
Technologies Used
Go (Golang): The programming language used to develop the API.
Gorilla Mux: A router and dispatcher for HTTP requests in Go.
Validator: A library for validating data like emails, age, and other constraints.
API Endpoints
Here are the available API endpoints:

POST /students: Create a new student.
GET /students: Fetch all students.
GET /students/{id}: Fetch a student by their ID.
PUT /students/{id}: Update student details.
DELETE /students/{id}: Delete a student by their ID.
