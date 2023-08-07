package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreateAt  time.Time          `bson:"create_at"`
	UpdateAt  time.Time          `bson:"update_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}
