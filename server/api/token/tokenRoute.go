package token

import (
	"net/http"
	"time"

	"github.com/onkarsutar/UserManagement/server/api/user"
	"github.com/onkarsutar/UserManagement/server/helper/logginghelper"
	"github.com/onkarsutar/UserManagement/server/model"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
)

func Init(o *echo.Group) {
	o.GET("/token/generate/:id", GenerateTokenRoute)
}

func GenerateTokenRoute(c echo.Context) error {
	id := c.Param("id")
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	userObj, err := user.GetUserByLoginIDDAO(id)
	claims["userName"] = userObj.UserName
	claims["id"] = userObj.LoginID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(model.JWTKey))
	if err != nil {
		logginghelper.LogError("GenerateTokenRoute Error: ", err)

		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
