package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "example.com/go-mongo-app/controllers"
    "example.com/go-mongo-app/services"
    "example.com/go-mongo-app/repositories"
)

func init() {
    godotenv.Load()
}
func main() {
    repo := repositories.NewStudentRepository()
    service := services.NewStudentService(repo)
    controller := controllers.NewStudentController(service)
    godotenv.Load()
    
    r := mux.NewRouter()

    r.HandleFunc("/students", controller.GetStudents).Methods("GET")
    r.HandleFunc("/students", controller.CreateStudent).Methods("POST")
    r.HandleFunc("/students/{id}", controller.GetStudentByID).Methods("GET")
    r.HandleFunc("/students/{id}", controller.DeleteStudent).Methods("DELETE")
    r.HandleFunc("/students/{id}", controller.UpdateStudent).Methods("PUT")

    fmt.Println("Servidor escuchando en el puerto 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}
