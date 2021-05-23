package model

import "time"

// User : Represents User Entity in Application.
type User struct {
	CreatedOn    time.Time `json:"createdOn" bson:"created_on"`
	ModifiedOn   time.Time `json:"modifiedOn" bson:"modified_on"`
	LoginID      string    `json:"loginID" bson:"login_id" validate:"required"`
	Password     string    `json:"password" bson:"password" validate:"required"`
	UserName     string    `json:"userName" bson:"user_name" validate:"required"`
	EmailID      string    `json:"emailID" bson:"email_id"`
	MobileNumber string    `json:"mobileNumber" bson:"mobile_number"`
}

type ChangedPassword struct {
	User
	NewPassword string `json:"newPassword" validate:"required"`
}
