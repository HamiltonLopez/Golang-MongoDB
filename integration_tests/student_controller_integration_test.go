package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/go-mongo-app/models"
	"example.com/go-mongo-app/repositories"
	"example.com/go-mongo-app/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *mux.Router {
	repo := repositories.NewStudentRepository()
	service := services.NewStudentService(repo)
	controller := NewStudentController(service)

	r := mux.NewRouter()
	r.HandleFunc("/students", controller.GetStudents).Methods("GET")
	r.HandleFunc("/students", controller.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", controller.GetStudentByID).Methods("GET")
	r.HandleFunc("/students/{id}", controller.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}", controller.UpdateStudent).Methods("PUT")


	return r
}

func TestCRUDIntegration(t *testing.T) {
	router := setupRouter()

	// 1. Crear estudiante
	student := models.Student{
		Name:  "Estudiante Prueba Hamilton",
		Age:   23,
		Email: "prueba@example.com",
	}
	body, _ := json.Marshal(student)

	req, _ := http.NewRequest("POST", "/students", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Obtener el ID del estudiante creado
	bodyResp, _ := ioutil.ReadAll(rr.Body)
	var createdStudent models.Student
	json.Unmarshal(bodyResp, &createdStudent)
	assert.NotEmpty(t, createdStudent.ID)

	// 2. Obtener todos los estudiantes
	reqGetAll, _ := http.NewRequest("GET", "/students", nil)
	rrGetAll := httptest.NewRecorder()
	router.ServeHTTP(rrGetAll, reqGetAll)
	assert.Equal(t, http.StatusOK, rrGetAll.Code)

	// 3. Obtener por ID
	urlByID := fmt.Sprintf("/students/%s", createdStudent.ID.Hex())
	reqGet, _ := http.NewRequest("GET", urlByID, nil)
	rrGet := httptest.NewRecorder()
	router.ServeHTTP(rrGet, reqGet)
	assert.Equal(t, http.StatusOK, rrGet.Code)

	// 4. Actualizar estudiante
	createdStudent.Name = "Nombre Actualizado"
updateBody, _ := json.Marshal(createdStudent)
reqUpdate, _ := http.NewRequest("PUT", "/students/"+createdStudent.ID.Hex(), bytes.NewBuffer(updateBody))
reqUpdate.Header.Set("Content-Type", "application/json")
rrUpdate := httptest.NewRecorder()
router.ServeHTTP(rrUpdate, reqUpdate)
assert.Equal(t, http.StatusOK, rrUpdate.Code) // ✅ ahora debería ser 200


	// 5. Eliminar estudiante
	reqDelete, _ := http.NewRequest("DELETE", urlByID, nil)
	rrDelete := httptest.NewRecorder()
	router.ServeHTTP(rrDelete, reqDelete)
	assert.Equal(t, http.StatusNoContent, rrDelete.Code)

}
