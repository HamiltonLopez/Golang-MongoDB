package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
    ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name  string             `bson:"name" json:"name"`
    Age   int                `bson:"age" json:"age"`
    Email string             `bson:"email" json:"email"`
}
