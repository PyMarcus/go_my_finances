package impl

import (
	"context"
	"log"

	"github.com/PyMarcus/go_finance/dao"
	"github.com/PyMarcus/go_finance/database"
	"github.com/PyMarcus/go_finance/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type financeImpl struct {
	finances []*models.Finance
}

func NewFinance() dao.Finance {
	return &financeImpl{}
}

func (f *financeImpl) GetAllRegisters() ([]*models.Finance, error) {

	conn := database.ConnectionFactory()
	defer conn.Disconnect(context.TODO())

	database := conn.Database("finance")
	collection := database.Collection("finances")

	cursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Panic(err)
		return nil, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result bson.M

		if err := cursor.Decode(&result); err != nil {
			log.Println(err)
			return nil, err
		}
		finance := models.NewFinance(result["_id"].(primitive.ObjectID),
			result["value"].(float64),
			result["name"].(string),
			result["category"].(string),
			result["date"].(string))

		f.finances = append(f.finances, finance)
	}

	return f.finances, nil
}

func (f *financeImpl) CreateRegister(finance *models.Finance) error {
	conn := database.ConnectionFactory()
	defer conn.Disconnect(context.TODO())

	database := conn.Database("finance")
	collection := database.Collection("finances")

	_, err := collection.InsertOne(context.TODO(), *finance)

	if err != nil{
		log.Panic(err)
		return err
	}

	log.Println("Created!")
	return nil
}

func (f *financeImpl) UpdateRegister(finance *models.Finance) error {
	return nil
}

func (f *financeImpl) DeleteRegister(id primitive.ObjectID) error {
	return nil
}
