package api

import (
	"net/http"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/onkarsutar/UserManagement/server/api/token"
	"github.com/onkarsutar/UserManagement/server/api/user"
	"github.com/onkarsutar/UserManagement/server/middleware"
	"github.com/onkarsutar/UserManagement/server/model"
)

func Init(e *echo.Echo) {

	o := e.Group("/o")
	r := e.Group("/r")
	// r.Use(echoMiddleware.JWTWithConfig(echoMiddleware.JWTConfig{
	// 	Claims:     &model.JwtCustomClaims{},
	// 	SigningKey: []byte(model.JWTKey),
	// }))

	r.Use(echoMiddleware.JWT([]byte(model.JWTKey)))

	middleware.Init(e, o, r)
	token.Init(o)
	user.Init(o, r)

	e.GET("/", RootRoute)
	e.GET("/checkststus", CheckStatusRoute)
}

// RootRoute : Entrypoint to API server.
func RootRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome")
}

// CheckStatusRoute : API to scheck server status.
func CheckStatusRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Running..")
}
