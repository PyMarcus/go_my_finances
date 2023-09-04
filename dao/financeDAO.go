package dao

import (
	"github.com/PyMarcus/go_finance/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Finance interface{
	GetAllRegisters() ([]*models.Finance, error)
	CreateRegister(finance *models.Finance) error
	UpdateRegister(finance *models.Finance) error 
	DeleteRegister(id primitive.ObjectID) error 
}
