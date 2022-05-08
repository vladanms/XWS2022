package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role int64

const (
	Admin   Role = 1
	RegUser      = 2
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Username string             `json:"username" validate:"username,excludesall= "`
	Email    string             `json:"email" validate:"required,excludesall= "`
	Password *string            `json:",omitempty" validate:"password,excludesall= "`
	Role     Role               `json:"-"`
	Public   bool               `json:"-"`
}

// Users defines a slice of Product
type Users []*User

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrUserNotFound = fmt.Errorf("User not found")

// GetUsers returns all users from the database
func GetUsers() Users {

	usersCollection := Client.Database("xws").Collection("users")

	// retrieve all the documents in a collection
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		panic(err)
	}

	var allUsers Users
	for cursor.Next(context.Background()) {
		tmp := User{}
		cursor.Decode(&tmp)
		tmpPtr := &tmp
		allUsers = append(allUsers, tmpPtr)
	}
	return allUsers

}

// GetUserByID returns a single user which matches the id from the
// database.
// If a User is not found this function returns a UserNotFound error
func GetUserByID(id string) (*User, error) {
	//loading db, setting context and converting id to a suitable parameter for FindOne() function
	usersCollection := Client.Database("xws").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("[ERROR] can't convert string to ObjectID", err)
		return nil, ErrUserNotFound
	}

	var result User
	err = usersCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)

	if err != nil {
		fmt.Println("[ERROR] FindOne() ObjectIDFromHex :", err)
		return nil, ErrUserNotFound
	}

	user := &result
	return user, nil
}

//All usernames are unique
func GetUserByUsername(u string) (*User, error) {
	usersCollection := Client.Database("xws").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	var result User
	filter := bson.D{{"username", u}}

	sRes := usersCollection.FindOne(ctx, filter)
	if sRes.Err() != nil {
		fmt.Println("[DEBUG] no user with that username")
		return nil, nil
	}
	err := sRes.Decode(&result)
	if err != nil {
		fmt.Println("[ERROR] could not decode result")
		return nil, err
	}
	fmt.Printf("\n%+v\n", result)
	user := &result
	return user, nil

}

// AddUser adds a new user to the database
func AddUser(u User) error {
	res, _ := GetUserByUsername(u.Username)
	if res != nil {
		return fmt.Errorf("user %s already registered", u.Username)
	}
	userCollection := Client.Database("xws").Collection("users")
	doc, err := bson.Marshal(u)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}
	result, err := userCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
	}
	fmt.Println(result.InsertedID)
	return nil
}
