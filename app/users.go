package pioj

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ObjectId primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username"`
	Password []byte             `json:"-"`
	Email    string             `json:"email"`
}

type UserWithPasswordAndVerification struct {
	User
	Password     string `json:"password"`
	Verification string `json:"verification"`
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

func (app App) HandleUsers() {
	app.ServeMux.HandleFunc("/api/user/create", func(w http.ResponseWriter, r *http.Request) {
		var user UserWithPasswordAndVerification
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := app.Redis.GetDel(context.TODO(), "pioj:verification:"+user.Email).Result()
		if err == redis.Nil {
			http.Error(w, "Verification code expired or not sent", http.StatusUnauthorized)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInsufficientStorage)
			return
		} else if res != user.Verification {
			http.Error(w, "Incorrect verification code", http.StatusUnauthorized)
			return
		}

		if err := user.User.SetPassword(user.Password); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := user.User.Create(context.TODO(), app.Database.Collection("users")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user.User)
	})

	app.ServeMux.HandleFunc("/api/user/email", func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		code := make([]byte, 6)
		if _, err := rand.Read(code); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for i := range code {
			code[i] = code[i]%10 + '0'
		}
		if err := app.Redis.Set(context.TODO(), "pioj:verification:"+email, code, 10*time.Minute).Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInsufficientStorage)
			return
		}
		if len(email) > 0 {
			err := SendMail(app.SMTP, []string{email}, "Email verification", string(code))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})
}
