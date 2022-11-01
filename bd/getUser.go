package bd

import (
	"context"
	"github.com/jorgeamc/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoC.Database("twittor")
	col := db.Collection("users")

	emailMatch := bson.M{"email": email}

	var user models.User

	err := col.FindOne(ctx, emailMatch).Decode(&user)
	ID := user.ID.Hex()
	if err != nil {
		return user, false, ID
	}
	return user, true, ID
}
