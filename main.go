
package main

import (
    "log"
    "net/http"
    "unique-student-api/routes"
)

func main() {
    r := routes.SetupRoutes()
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
