package main

import (
	"fmt"
	"log"

	"github.com/PyMarcus/go_finance/controllers"
	"github.com/PyMarcus/go_finance/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	data, err := controllers.GetAllRegisters()

	if err != nil {
		log.Panic(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}


	objectID:= primitive.NewObjectID()

	f := models.Finance{
		Id:    objectID,
		Name:  "ij",
		Value: -232.54,
		Date: "2023-09-03 23:32:33",
		Category: "CONTA A PAGAR",
	}

	controllers.CreateRegister(&f)
}
