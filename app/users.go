package pioj

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ObjectId primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username"`
	Password []byte             `json:"password"`
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

func (user *User) Create(ctx context.Context, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	user.ObjectId = primitive.NewObjectID()
	return collection.InsertOne(
		ctx,
		user,
	)
}

func (user *User) Update(ctx context.Context, collection *mongo.Collection) (*mongo.UpdateResult, error) {
	return collection.ReplaceOne(ctx, map[string]any{"_id": user.ObjectId}, user)
}

func (user *User) Save(ctx context.Context, collection *mongo.Collection) (any, error) {
	if indexName, err := SetUnique(ctx, collection, "username"); err != nil {
		return indexName, err
	}
	if user.ObjectId == primitive.NilObjectID {
		return user.Create(ctx, collection)
	} else {
		return user.Update(ctx, collection)
	}
}
