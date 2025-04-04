package controllers

import (
    "encoding/json"
    "net/http"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/services"
)

type StudentController struct {
    service *services.StudentService
}

func NewStudentController(service *services.StudentService) *StudentController {
    return &StudentController{service}
}

func (c *StudentController) GetStudents(w http.ResponseWriter, r *http.Request) {
    students, err := c.service.GetStudents()
    if err != nil {
        http.Error(w, "Error al obtener estudiantes", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(students)
}

func (c *StudentController) CreateStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student
    err := json.NewDecoder(r.Body).Decode(&student)
    if err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    newStudent, err := c.service.AddStudent(student)
    if err != nil {
        http.Error(w, "Error al insertar estudiante", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(newStudent)
}
