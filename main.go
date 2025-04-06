package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "example.com/go-mongo-app/controllers"
    "example.com/go-mongo-app/services"
    "example.com/go-mongo-app/repositories"
)

func main() {
    repo := repositories.NewStudentRepository()
    service := services.NewStudentService(repo)
    controller := controllers.NewStudentController(service)

    r := mux.NewRouter()

    r.HandleFunc("/students", controller.GetStudents).Methods("GET")
    r.HandleFunc("/students", controller.CreateStudent).Methods("POST")
    r.HandleFunc("/students/{id}", controller.GetStudentByID).Methods("GET")
    r.HandleFunc("/students/{id}", controller.DeleteStudent).Methods("DELETE")
    r.HandleFunc("/students", controller.UpdateStudent).Methods("PUT")

    fmt.Println("Servidor escuchando en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}
