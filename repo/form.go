package repo

import (
	"context"
	"errors"
	"portfolio_form/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FormRepo struct {
	mongoDb *mongo.Database
}

func InitializeFormRepo(mongoDb *mongo.Database) *FormRepo {
	return &FormRepo{
		mongoDb: mongoDb,
	}
}

func (f *FormRepo) CreateFormDetails(payload models.FormInputs) error {
	payload.CreatedAt = time.Now()
	database := f.mongoDb.Collection("forms")
	_, err := database.InsertOne(context.Background(), payload)
	if err != nil {
		return err
	}
	return nil
}

func (f *FormRepo) GetAllFormDetailsByEmail(email string) ([]models.FormInputs, error) {
	collection := f.mongoDb.Collection("forms")
	filter := bson.M{"ownermail": email}
	records, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, errors.New("error in collections")
	}
	var formInputs []models.FormInputs
	if err := records.All(context.Background(), &formInputs); err != nil {
		return nil, errors.New("error in fetching data")
	}

	return formInputs, nil
}
