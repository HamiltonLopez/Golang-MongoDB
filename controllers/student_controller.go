package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/services"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentController struct {
    Service services.StudentServiceInterface
}
func NewStudentController(service services.StudentServiceInterface) *StudentController {
    return &StudentController{
        Service: service,
    }
}

func (c *StudentController) GetStudents(w http.ResponseWriter, r *http.Request) {
    students, err := c.Service.GetStudents()
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

    newStudent, err := c.Service.AddStudent(student)
    if err != nil {
        http.Error(w, "Error al insertar estudiante", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newStudent)
    
}

func (c *StudentController) GetStudentByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "ID no proporcionado", http.StatusBadRequest)
        return
    }

    student, err := c.Service.GetStudentByID(id)
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
    vars := mux.Vars(r) // 👈 extraemos el id desde la URL
    id := vars["id"]

    var student models.Student
    err := json.NewDecoder(r.Body).Decode(&student)
    if err != nil {
        http.Error(w, "Datos inválidos", http.StatusBadRequest)
        return
    }

    // Convertimos el ID a ObjectID y lo asignamos al struct
    student.ID, err = primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, "ID inválido", http.StatusBadRequest)
        return
    }

    updatedStudent, err := c.Service.UpdateStudent(&student)
    if err != nil {
        http.Error(w, "Error al actualizar estudiante", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
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

    err := c.Service.DeleteStudentByID(id)
    if err != nil {
        http.Error(w, "Error al eliminar estudiante", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}

