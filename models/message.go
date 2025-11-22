package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type message struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	GroupID   *primitive.ObjectID  `bson:"group_id,omitempty" json:"group_id,omitempty"`
	UserID    primitive.ObjectID   `bson:"user_id" json:"user_id"`
	Content   string               `bson:"content" json:"content"`
	Type      string               `bson:"type" json:"type"`
	Timestamp int64                `bson:"timestamp" json:"timestamp"`
	SeenBy    []primitive.ObjectID `bson:"seen_by,omitempty" json:"seen_by,omitempty"`
	Reaction  string               `bson:"reaction,omitempty" json:"reaction,omitempty"`
	Replies   []primitive.ObjectID `bson:"replies,omitempty" json:"replies,omitempty"`
}