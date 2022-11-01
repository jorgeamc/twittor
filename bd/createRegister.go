package bd

import (
	"github.com/jorgeamc/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	"time"
)

func CreateRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("twittor")
	col := db.Collection("users")

	password, err := EncryptPassword(u.Password)
	if err != nil {
		return "", false, err
	}
	u.Password = password

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return id.String(), true, nil
}
