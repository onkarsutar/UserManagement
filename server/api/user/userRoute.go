package user

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/onkarsutar/UserManagement/server/helper/logginghelper"
	"github.com/onkarsutar/UserManagement/server/model"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// Init : Inits API
func Init(o, r *echo.Group) {
	validate = validator.New()
	o.POST("/user/login", LoginRoute)
	o.POST("/user/changepassword", ChangePasswordRoute)
	o.POST("/user/add", AddUserRoute)
}

// LoginRoute : API to validate user credentials.
func LoginRoute(c echo.Context) error {
	userObj := model.User{}
	err := c.Bind(&userObj)
	if err != nil {
		logginghelper.LogError("LoginRoute Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	err = ValidateLoginService(userObj)
	if err != nil {
		logginghelper.LogError("LoginRoute Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, "valid user.")
}

// AddUserRoute : API to create new user.
func AddUserRoute(c echo.Context) error {
	userObj := model.User{}
	err := c.Bind(&userObj)
	if err != nil {
		logginghelper.LogError("AddUserRoute Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	err = validate.Struct(userObj)
	if err != nil {
		logginghelper.LogError("AddUserRoute Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	err = AddUserService(userObj)
	if err != nil {
		logginghelper.LogError("AddUserRoute Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, "User Added.")
}

// ChangePasswordRoute : API to change password of user.
func ChangePasswordRoute(c echo.Context) error {
	changePasswordObj := model.ChangedPassword{}
	err := c.Bind(&changePasswordObj)
	if err != nil {
		logginghelper.LogError("ChangePasswordRoute Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	// err = validate.Struct(changePasswordObj)
	// if err != nil {
	// 	logginghelper.LogError("ChangePasswordRoute Error: ", err)
	// 	return c.JSON(http.StatusExpectationFailed, err.Error())
	// }
	err = ChangePasswordService(changePasswordObj)
	if err != nil {
		logginghelper.LogError("ChangePasswordRoute Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, "Password Changed.")
}
