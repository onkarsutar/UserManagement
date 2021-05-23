package user

import (
	"errors"
	"unicode"

	"github.com/onkarsutar/UserManagement/server/helper/logginghelper"
	"github.com/onkarsutar/UserManagement/server/model"
)

func ValidateLoginService(userObj model.User) error {
	dbUser, err := GetUserByLoginIDDAO(userObj.LoginID)
	if err != nil {
		logginghelper.LogError("ValidateLoginService Error: ", err)
		return err
	}
	if dbUser.Password != userObj.Password {
		logginghelper.LogError("ValidateLoginService Error: Password Not Matching.")
		return errors.New("password not matching")
	}
	return nil
}

func AddUserService(userObj model.User) error {
	status := PasswordValidator(userObj.Password)
	if !status {
		logginghelper.LogError("AddUserService Error : Not Valid Password.")
		return errors.New("not valid password")
	}
	isExists, err := IsUserExists(userObj)
	if err != nil {
		logginghelper.LogError("AddUserService Error: ", err)
		return err
	}
	if isExists {
		logginghelper.LogError("AddUserService Error : User Already Exists.")
		return errors.New("user already exists")
	}
	err = AddUserDAO(userObj)
	return err
}

func ChangePasswordService(changePasswordObj model.ChangedPassword) error {
	if changePasswordObj.NewPassword == changePasswordObj.Password {
		logginghelper.LogError("ChangePasswordService Error : New Password & Old Password Are Not Same.")
		return errors.New("old password and new password are not same.")

	}
	status := PasswordValidator(changePasswordObj.NewPassword)
	if !status {
		logginghelper.LogError("ChangePasswordService Error : Not Valid Password.")
		return errors.New("not valid password")
	}
	isValid, err := ValidateUser(changePasswordObj.User)
	if err != nil {
		logginghelper.LogError("ChangePasswordService Error: ", err)
		return err
	}
	if !isValid {
		logginghelper.LogError("ChangePasswordService Error : User Not Valid.")
		return errors.New("user not valid")
	}
	err = ChangePasswordDAO(changePasswordObj)
	return err
}

func IsUserExists(userObj model.User) (bool, error) {
	dbUser, err := GetUserByLoginIDDAO(userObj.LoginID)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return false, nil
	}
	if err != nil {
		logginghelper.LogError("IsUserExists Error: ", err)
		return true, err
	}
	// if dbUser.LoginID == userObj.LoginID {
	// 	logginghelper.LogError("IsUserExists Error: User Already Exists.")
	// 	return true, nil //errors.New("user already exists")
	// }
	logginghelper.LogError("IsUserExists Error: User Already Exists. ", dbUser.LoginID)
	return true, nil
}

func ValidateUser(userObj model.User) (bool, error) {
	dbUser, err := GetUserByLoginIDDAO(userObj.LoginID)
	if err != nil && err.Error() == "mongo: no documents in result" {
		return false, nil
	}
	if err != nil {
		logginghelper.LogError("ValidateUser Error: ", err)
		return false, err
	}
	if dbUser.Password != userObj.Password {
		logginghelper.LogError("ValidateUser Error: Password Not Matched.")
		return false, nil //errors.New("user already exists")
	}
	return true, nil
}

func PasswordValidator(password string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}
