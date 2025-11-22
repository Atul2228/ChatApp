package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type user struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string               `bson:"email" json:"email"`
	OTP       string               `bson:"otp,omitempty" json:"otp,omitempty"`
	SocketID  string               `bson:"socket_id,omitempty" json:"socket_id,omitempty"`
	ServerID  string               `bson:"server_id,omitempty" json:"server_id,omitempty"`
	UserImage string               `bson:"user_image,omitempty" json:"user_image,omitempty"`
	Messages  []primitive.ObjectID `bson:"messages,omitempty" json:"messages,omitempty"`
}