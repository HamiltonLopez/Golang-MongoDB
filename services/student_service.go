package services

import (
    "example.com/go-mongo-app/models"
    "example.com/go-mongo-app/repositories"
)

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
