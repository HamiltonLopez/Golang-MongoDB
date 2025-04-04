package repositories

import (
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "example.com/go-mongo-app/models"
    "log"
)

type StudentRepository struct {
    collection *mongo.Collection
}

func NewStudentRepository() *StudentRepository {
    clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    collection := client.Database("school").Collection("students")
    return &StudentRepository{collection}
}

func (repo *StudentRepository) GetAllStudents() ([]models.Student, error) {
    var students []models.Student
    cursor, err := repo.collection.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var student models.Student
        if err := cursor.Decode(&student); err != nil {
            return nil, err
        }
        students = append(students, student)
    }

    return students, nil
}

func (repo *StudentRepository) CreateStudent(student models.Student) (*models.Student, error) {
    student.ID = primitive.NewObjectID()
    _, err := repo.collection.InsertOne(context.TODO(), student)
    if err != nil {
        return nil, err
    }
    return &student, nil
}
