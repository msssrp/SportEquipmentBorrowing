package user

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userRepositoryMongo struct {
	collection *mongo.Collection
	secretKey  []byte
}

func NewUserRepositoryMongo(client *mongo.Client, dbName string, collectionName string, secretKey []byte) UserRepository {
	collection := client.Database(dbName).Collection(collectionName)
	return &userRepositoryMongo{
		collection: collection,
		secretKey:  []byte(secretKey),
	}
}

//Get
func (r *userRepositoryMongo) GetAll() ([]*User, error) {
	var users []*User

	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepositoryMongo) GetByID(id primitive.ObjectID) (*User, error) {
	filter := bson.M{"_id": id}

	var user User
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

//Post
func (r *userRepositoryMongo) Create(user *User) error {
	var existingUser User

	err := r.collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)

	if err == mongo.ErrNoDocuments {
		// Username is not taken, proceed with creating the user
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)

		_, err = r.collection.InsertOne(context.Background(), user)
		return err
	} else if err != nil {
		// Some other error occurred
		return err
	}

	// Username is already taken
	return errors.New("username already taken")
}

func (r *userRepositoryMongo) SignIn(username string, password string) (string, string, error) {
	var user User
	var errInvalidCredentials = errors.New("Invalid username or password")
	filter := bson.M{"username": username}
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// User not found
		return "", "", errInvalidCredentials
	}
	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// Passwords do not match
		return "", "", errInvalidCredentials
	}

	// Generate JWT token
	accessToken, refreshToken, err := generateJWTToken(user, []byte(r.secretKey))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (r *userRepositoryMongo) GenerateNewAccessToken(userID string) (string, error) {
	newAccessToken, err := generateNewAccessToken(userID, r.secretKey)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}

//Put
func (r *userRepositoryMongo) Update(user *User) error {
	filter := bson.M{"_id": user.Id}
	update := bson.M{
		"$set": user,
	}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	return err
}

//Delete
func (r *userRepositoryMongo) DeleteByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

//func
// generateJWTToken generates a JWT token for the user.
func generateJWTToken(user User, secretKey []byte) (string, string, error) {
	accessClaims := jwt.MapClaims{
		"user_id": user.Id.Hex(),
		// Add more claims as needed
		"exp": time.Now().Add(time.Minute * 30).Unix(), // Access token expires in 30 minutes
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	signedAccessToken, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	// Refresh token
	refreshClaims := jwt.MapClaims{
		"user_id": user.Id.Hex(),
		"exp":     time.Now().Add(time.Hour * 12).Unix(), // Refresh token expires in 12 hours
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedRefreshToken, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", err
	}

	return signedAccessToken, signedRefreshToken, nil
}

func generateNewAccessToken(userID string, secretKey []byte) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 30).Unix(), // Token expires in 30 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newAccesToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return newAccesToken, nil
}
