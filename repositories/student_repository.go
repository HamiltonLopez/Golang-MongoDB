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

func (repo *StudentRepository) GetStudentByID(id string) (*models.Student, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var student models.Student
    filter := bson.M{"_id": objectID}
    err = repo.collection.FindOne(context.TODO(), filter).Decode(&student)
    if err != nil {
        return nil, err
    }

    return &student, nil
}

func (repo *StudentRepository) UpdateStudent(student models.Student) (*models.Student, error) {
    filter := bson.M{"_id": student.ID}
    update := bson.M{"$set": student}

    _, err := repo.collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return nil, err
    }

    return &student, nil
}

// repositories/student_repository.go
func (r *StudentRepository) RemoveStudentByID(id string) (bool, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return false, err
    }

    result, err := r.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        return false, err
    }

    return result.DeletedCount > 0, nil
}
