package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Finance struct {
	Id       primitive.ObjectID     `json:"_id" bson:"_id"`
	Name     string  `json:"name" bson:"name"`
	Value    float64 `json:"value" bson:"value"`
	Category string  `json:"category" bson:"category"`
	Date     string  `json:"date" bson:"date"`
}

func NewFinance(id primitive.ObjectID, value float64, name, category, date string) *Finance {
	return &Finance{
		Id:       id,
		Name:     name,
		Value:    value,
		Category: category,
		Date:     date,
	}
}
