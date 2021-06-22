package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* BorroTweet borra un Tweet determinado */
func BorroTweet(ID string, UserId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	bd := MongoCN.Database("twittor")
	col := bd.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id":    objID,
		"userid": UserId,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}
