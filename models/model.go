package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bot struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Os       string             `json:"os,omitempty" validate:"required"`
	Ip       string             `json:"ip,omitempty" validate:"required"`
	Location string             `json:"location,omitempty" validate:"required"`
}
