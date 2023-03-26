package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Nextflix struct {
	ID        primitive.ObjectID `json:"_id,omitempty"         bson:"_id,omitempty"`
	Moviename string             `json:"movie,omitempty"`
	Watched   bool               `json:"watched,omitempty"`
}
