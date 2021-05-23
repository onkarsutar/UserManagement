package user

import (
	"context"
	"time"

	"github.com/onkarsutar/UserManagement/server/helper/confighelper"
	"github.com/onkarsutar/UserManagement/server/helper/dalhelper"
	"github.com/onkarsutar/UserManagement/server/helper/logginghelper"
	"github.com/onkarsutar/UserManagement/server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUserByLoginIDDAO : Get user from DB by loginID
func GetUserByLoginIDDAO(loginID string) (model.User, error) {

	userObj := model.User{}
	filter := bson.D{primitive.E{Key: "login_id", Value: loginID}}

	session, err := dalhelper.GetMongoClient()
	if err != nil {
		logginghelper.LogError("GetUserByLoginIDDAO Error: ", err)
		return userObj, err
	}
	collection := session.Database(confighelper.GetConfig("db")).Collection(model.USER_COLLECTION)
	err = collection.FindOne(context.TODO(), filter).Decode(&userObj)
	if err != nil {
		logginghelper.LogError("GetUserByLoginIDDAO Error: ", err)
		return userObj, err
	}

	return userObj, nil
}

// AddUserDAO : Add new user in DB
func AddUserDAO(userObj model.User) error {
	userObj.CreatedOn = time.Now()
	session, err := dalhelper.GetMongoClient()
	if err != nil {
		logginghelper.LogError("AddUser Error: ", err)
		return err
	}
	collection := session.Database(confighelper.GetConfig("db")).Collection(model.USER_COLLECTION)
	_, err = collection.InsertOne(context.TODO(), userObj)
	if err != nil {
		logginghelper.LogError("AddUser Error : ", err)
		return err
	}
	return nil
}

func ChangePasswordDAO(changePasswordObj model.ChangedPassword) error {
	changePasswordObj.ModifiedOn = time.Now()
	filter := bson.D{primitive.E{Key: "login_id", Value: changePasswordObj.LoginID}}

	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "password", Value: changePasswordObj.NewPassword},
	}}}

	session, err := dalhelper.GetMongoClient()
	if err != nil {
		logginghelper.LogError("ChangePasswordDAO Error: ", err)
		return err
	}
	collection := session.Database(confighelper.GetConfig("db")).Collection(model.USER_COLLECTION)
	_, err = collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		logginghelper.LogError("ChangePasswordDAO Error : ", err)
		return err
	}
	return nil
}
