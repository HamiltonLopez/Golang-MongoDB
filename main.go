package main

import (
    "fmt"
    "log"
    "net/http"
    "example.com/go-mongo-app/controllers"
    "example.com/go-mongo-app/services"
    "example.com/go-mongo-app/repositories"
)

func main() {
    repo := repositories.NewStudentRepository()
    service := services.NewStudentService(repo)
    controller := controllers.NewStudentController(service)

    http.HandleFunc("/students", controller.GetStudents)
    http.HandleFunc("/students/create", controller.CreateStudent)

    fmt.Println("Servidor escuchando en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
