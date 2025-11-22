package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Group struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string               `bson:"name" json:"name"`
	UserIDs []primitive.ObjectID `bson:"user_ids" json:"user_ids"`
}