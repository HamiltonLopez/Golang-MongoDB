package controllers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/go-mongo-app/models"
)

// ---------------------- MOCK DEL SERVICIO ----------------------

type MockStudentService struct {
    mock.Mock
}



func (m *MockStudentService) GetStudents() ([]models.Student, error) {
	args := m.Called()
	return args.Get(0).([]models.Student), args.Error(1)
}

func (m *MockStudentService) GetStudentByID(id string) (*models.Student, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockStudentService) AddStudent(student models.Student) (*models.Student, error) {
	args := m.Called(student)
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockStudentService) UpdateStudent(student *models.Student) (*models.Student, error) {
    args := m.Called(student)
    return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockStudentService) DeleteStudentByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}


// ---------------------- TESTS ----------------------

func TestCreateStudent_Success(t *testing.T) {
    mockService := new(MockStudentService)
    controller := NewStudentController(mockService)

    student := models.Student{
        Name:  "Juan",
        Age:   22,
        Email: "juan@test.com",
    }

    body, _ := json.Marshal(student)
    req := httptest.NewRequest("POST", "/students", bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()

    mockService.On("AddStudent", mock.MatchedBy(func(s models.Student) bool {
        return s.Name == student.Name && s.Age == student.Age && s.Email == student.Email
    })).Return(&models.Student{
        ID:    primitive.NewObjectID(),
        Name:  student.Name,
        Age:   student.Age,
        Email: student.Email,
    }, nil)
    
    controller.CreateStudent(resp, req)

    assert.Equal(t, http.StatusCreated, resp.Code)
    mockService.AssertExpectations(t)
}


func TestGetStudents_Success(t *testing.T) {
    mockService := new(MockStudentService)
    controller := NewStudentController(mockService)

    students := []models.Student{
        {Name: "Ana", Age: 21, Email: "ana@test.com"},
        {Name: "Luis", Age: 23, Email: "luis@test.com"},
    }

    mockService.On("GetStudents").Return(students, nil)

    req := httptest.NewRequest("GET", "/students", nil)
    resp := httptest.NewRecorder()

    controller.GetStudents(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    mockService.AssertExpectations(t)
}

func TestGetStudentByID_Success(t *testing.T) {
    mockService := new(MockStudentService)
    controller := NewStudentController(mockService)

    id := primitive.NewObjectID().Hex()
    student := &models.Student{ID: primitive.NewObjectID(), Name: "Carlos", Age: 25, Email: "carlos@test.com"}

    mockService.On("GetStudentByID", id).Return(student, nil)

    req := httptest.NewRequest("GET", "/students/"+id, nil)
    req = mux.SetURLVars(req, map[string]string{"id": id})
    resp := httptest.NewRecorder()

    controller.GetStudentByID(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    mockService.AssertExpectations(t)
}

func TestDeleteStudent_Success(t *testing.T) {
    mockService := new(MockStudentService)
    controller := NewStudentController(mockService)

    id := primitive.NewObjectID().Hex()

    mockService.On("DeleteStudentByID", mock.Anything).Return(nil)


    req := httptest.NewRequest("DELETE", "/students/"+id, nil)
    req = mux.SetURLVars(req, map[string]string{"id": id})
    resp := httptest.NewRecorder()

    controller.DeleteStudent(resp, req)

    assert.Equal(t, http.StatusNoContent, resp.Code)
    mockService.AssertExpectations(t)
}
func TestUpdateStudent_Success(t *testing.T) {
    mockService := new(MockStudentService)
    controller := NewStudentController(mockService)

    id := primitive.NewObjectID().Hex()
    updatedStudent := &models.Student{
        ID:    primitive.NewObjectID(),
        Name:  "Updated Name",
        Age:   30,
        Email: "updated@test.com",
    }

    body, _ := json.Marshal(updatedStudent)
    req := httptest.NewRequest("PUT", "/students/"+id, bytes.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    req = mux.SetURLVars(req, map[string]string{"id": id})
    resp := httptest.NewRecorder()

    mockService.On("UpdateStudent", mock.MatchedBy(func(s *models.Student) bool {
        return s.Name == updatedStudent.Name && s.Age == updatedStudent.Age && s.Email == updatedStudent.Email
    })).Return(updatedStudent, nil)

    controller.UpdateStudent(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)
    mockService.AssertExpectations(t)
}
