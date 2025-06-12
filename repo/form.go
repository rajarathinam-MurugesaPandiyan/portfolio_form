package repo

import (
	"context"
	"portfolio_form/models"
	"time"

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
