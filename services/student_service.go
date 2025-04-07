package services

import (
    "fmt"
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/repositories"
)
type StudentServiceInterface interface {
	GetStudents() ([]models.Student, error)
	GetStudentByID(id string) (*models.Student, error)
	AddStudent(student models.Student) (*models.Student, error)
	UpdateStudent(student *models.Student) (*models.Student, error)
	DeleteStudentByID(id string) error
}



type StudentService struct {
    repo *repositories.StudentRepository
}

func NewStudentService(repo *repositories.StudentRepository) *StudentService {
    return &StudentService{repo}
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
    return s.repo.GetAllStudents()
}

func (s *StudentService) AddStudent(student models.Student) (*models.Student, error) {
    return s.repo.CreateStudent(student)
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
    return s.repo.GetStudentByID(id)
}

// services/student_service.go
func (s *StudentService) DeleteStudentByID(id string) error {
    deleted, err := s.repo.RemoveStudentByID(id)
    if err != nil {
        return err
    }
    if !deleted {
        return fmt.Errorf("estudiante no encontrado")
    }
    return nil
}


func (s *StudentService) UpdateStudent(student *models.Student) (*models.Student, error) {
    return s.repo.UpdateStudent(student)
}
