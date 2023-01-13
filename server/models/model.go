package models

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Bot struct {
	Os       string `json:"os,omitempty" validate:"required"`
	Ram      string `json:"ram,omitempty" validate:"required"`
	Cpu      string `json:"cpu,omitempty" validate:"required"`
	Disk     string `json:"disk,omitempty" validate:"required"`
	Platform string `json:"platform,omitempty" validate:"required"`
	Hostname string `json:"hostname,omitempty" validate:"required"`
	Ip       string `json:"ip,omitempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
}
