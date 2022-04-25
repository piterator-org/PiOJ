package pioj

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
}

func (user *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err == nil {
		user.Password = hash
	}
	return err
}

func (user *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword(
		user.Password,
		[]byte(password),
	) == nil
}

func (user *User) Save(collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(
		context.TODO(),
		user,
	)
}
