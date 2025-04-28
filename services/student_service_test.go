package services

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "fmt"

    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson/primitive"

	"example.com/go-mongo-app/models"
	"example.com/go-mongo-app/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockStudentRepository struct {
	mock.Mock
}

func (m *MockStudentRepository) GetAllStudents() ([]models.Student, error) {
	args := m.Called()
	return args.Get(0).([]models.Student), args.Error(1)
}

func (m *MockStudentRepository) CreateStudent(student models.Student) (*models.Student, error) {
	args := m.Called(student)
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockStudentRepository) GetStudentByID(id string) (*models.Student, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Student), args.Error(1)
}

func (m *MockStudentRepository) RemoveStudentByID(id string) (bool, error) {
	args := m.Called(id)
	return args.Bool(0), args.Error(1)
}

func (m *MockStudentRepository) UpdateStudent(student *models.Student) (*models.Student, error) {
	args := m.Called(student)
	return args.Get(0).(*models.Student), args.Error(1)
}
func TestGetStudents(t *testing.T) {
    mockRepo := new(MockStudentRepository)
    studentList := []models.Student{
        {Name: "Ana", Age: 22, Email: "ana@test.com"},
    }

    mockRepo.On("GetAllStudents").Return(studentList, nil)

    service := services.NewStudentService(mockRepo)
    result, err := service.GetStudents()

    assert.NoError(t, err)
    assert.Equal(t, studentList, result)
    mockRepo.AssertExpectations(t)
}

func TestGetStudentByID_Success(t *testing.T) {
    mockRepo := new(MockStudentRepository)
    student := &models.Student{ID: "123", Name: "Pedro", Age: 20, Email: "pedro@test.com"}

    mockRepo.On("GetStudentByID", "123").Return(student, nil)

    service := services.NewStudentService(mockRepo)
    result, err := service.GetStudentByID("123")

    assert.NoError(t, err)
    assert.Equal(t, student, result)
    mockRepo.AssertExpectations(t)
}

func TestAddStudent(t *testing.T) {
    mockRepo := new(MockStudentRepository)
    student := models.Student{Name: "Luis", Age: 21, Email: "luis@test.com"}

    mockRepo.On("CreateStudent", student).Return(&student, nil)

    service := services.NewStudentService(mockRepo)
    result, err := service.AddStudent(student)

    assert.NoError(t, err)
    assert.Equal(t, &student, result)
    mockRepo.AssertExpectations(t)
}

func TestUpdateStudent(t *testing.T) {
    mockRepo := new(MockStudentRepository)
    student := &models.Student{ID: "123", Name: "Editado", Age: 25, Email: "editado@test.com"}

    mockRepo.On("UpdateStudent", student).Return(student, nil)

    service := services.NewStudentService(mockRepo)
    result, err := service.UpdateStudent(student)

    assert.NoError(t, err)
    assert.Equal(t, student, result)
    mockRepo.AssertExpectations(t)
}

func TestDeleteStudentByID_Success(t *testing.T) {
    mockRepo := new(MockStudentRepository)

    mockRepo.On("RemoveStudentByID", "123").Return(true, nil)

    service := services.NewStudentService(mockRepo)
    err := service.DeleteStudentByID("123")

    assert.NoError(t, err)
    mockRepo.AssertExpectations(t)
}

func TestDeleteStudentByID_NotFound(t *testing.T) {
    mockRepo := new(MockStudentRepository)

    mockRepo.On("RemoveStudentByID", "notfound").Return(false, nil)

    service := services.NewStudentService(mockRepo)
    err := service.DeleteStudentByID("notfound")

    assert.Error(t, err)
    assert.EqualError(t, err, "estudiante no encontrado")
    mockRepo.AssertExpectations(t)
}