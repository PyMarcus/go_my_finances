package controllers

import (
	"github.com/PyMarcus/go_finance/dao/impl"
	"github.com/PyMarcus/go_finance/models"
)

func GetAllRegisters() ([]*models.Finance, error){
	finance := impl.NewFinance()
	return finance.GetAllRegisters()
}

func CreateRegister(finance *models.Finance) error{
	financeStruct := impl.NewFinance()
	return financeStruct.CreateRegister(finance)
}