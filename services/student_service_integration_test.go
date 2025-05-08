package services

import (
    "context"
    "testing"
    "time"

    "example.com/go-mongo-app/models"
	"example.com/go-mongo-app/database"
    "example.com/go-mongo-app/repositories"

    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/bson"
)

func TestStudentServiceIntegration(t *testing.T) {
	// Setup database connection
	db, err := database.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Client().Disconnect(context.Background())

	studentRepo := repositories.NewStudentRepository(db)
	studentService := NewStudentService(studentRepo)

	// Clean up the collection before and after the test
	collection := db.Collection("students")
	defer collection.Drop(context.Background())

	// Setup router
	router := setupRouter(studentService)

	t.Run("CreateStudent", func(t *testing.T) {
		student := models.Student{
			Name:  "John Doe",
			Age:   21,
			Email: "johndoe@example.com",
		}

		createdStudent, err := studentService.CreateStudent(context.Background(), student)
		assert.NoError(t, err)
		assert.NotNil(t, createdStudent.ID)

		var result models.Student
		err = collection.FindOne(context.Background(), bson.M{"_id": createdStudent.ID}).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, student.Name, result.Name)
		assert.Equal(t, student.Age, result.Age)
		assert.Equal(t, student.Email, result.Email)
	})

	t.Run("GetStudentByID", func(t *testing.T) {
		student := models.Student{
			Name:  "Jane Doe",
			Age:   22,
			Email: "janedoe@example.com",
		}

		createdStudent, err := studentService.CreateStudent(context.Background(), student)
		assert.NoError(t, err)

		fetchedStudent, err := studentService.GetStudentByID(context.Background(), createdStudent.ID)
		assert.NoError(t, err)
		assert.Equal(t, student.Name, fetchedStudent.Name)
		assert.Equal(t, student.Age, fetchedStudent.Age)
		assert.Equal(t, student.Email, fetchedStudent.Email)
	})

	t.Run("UpdateStudent", func(t *testing.T) {
		student := models.Student{
			Name:  "Alice",
			Age:   23,
			Email: "alice@example.com",
		}

		createdStudent, err := studentService.CreateStudent(context.Background(), student)
		assert.NoError(t, err)

		updatedData := models.Student{
			Name:  "Alice Updated",
			Age:   24,
			Email: "alice.updated@example.com",
		}

		err = studentService.UpdateStudent(context.Background(), createdStudent.ID, updatedData)
		assert.NoError(t, err)

		var result models.Student
		err = collection.FindOne(context.Background(), bson.M{"_id": createdStudent.ID}).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, updatedData.Name, result.Name)
		assert.Equal(t, updatedData.Age, result.Age)
		assert.Equal(t, updatedData.Email, result.Email)
	})
}

func setupRouter(studentService *StudentService) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/students", studentService.HandleCreateStudent)
	router.HandleFunc("/students/{id}", studentService.HandleGetStudentByID)
	router.HandleFunc("/students/{id}/update", studentService.HandleUpdateStudent)
	return router
}