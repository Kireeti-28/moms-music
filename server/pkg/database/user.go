package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	databaseName   = "tidings"
	collectionName = "users"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDB struct {
	ID           string `bson:"_id"`
	Email        string `bson:"email"`
	HashPassword string `bson:"password"`
}

func InsertUser(user User) error {
	client, err := getClient()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())

	_, err = client.Database(databaseName).Collection(collectionName).InsertOne(context.TODO(), bson.D{
		{Key: "email", Value: user.Email},
		{Key: "password", Value: user.Password},
	})
	if err != nil {
		return err
	}

	return nil
}

func GetUser(email string) (UserDB, error) {
	client, err := getClient()
	if err != nil {
		return UserDB{}, err
	}
	defer client.Disconnect(context.TODO())

	userDB := UserDB{}
	err = client.Database(databaseName).Collection(collectionName).FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&userDB)
	if err != nil {
		return UserDB{}, err
	}

	return userDB, nil
}
