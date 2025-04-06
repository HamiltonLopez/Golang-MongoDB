package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
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

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message":  "Estudiantes obtenidos correctamente",
        "students": students,
    })
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

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Estudiante creado correctamente",
        "student": newStudent,
    })
}

func (c *StudentController) GetStudentByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }

    student, err := c.service.GetStudentByID(id)
    if err != nil {
        http.Error(w, "Estudiante no encontrado", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Estudiante encontrado",
        "student": student,
    })
}

func (c *StudentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student
    err := json.NewDecoder(r.Body).Decode(&student)
    if err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    updatedStudent, err := c.service.UpdateStudent(student)
    if err != nil {
        http.Error(w, "Error al actualizar estudiante", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Estudiante actualizado correctamente",
        "student": updatedStudent,
    })
}

func (c *StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }

    err := c.service.DeleteStudentByID(id)
    if err != nil {
        http.Error(w, "Error al eliminar estudiante", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Estudiante eliminado correctamente",
    })
}
