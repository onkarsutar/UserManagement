package main

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/onkarsutar/UserManagement/server/api"
	"github.com/onkarsutar/UserManagement/server/helper/confighelper"
	"github.com/onkarsutar/UserManagement/server/helper/logginghelper"
)

var (
	apiServerPort = "4365"

	maxBackupCount       = 20
	maxAgeForBackupFiles = 20
	maxBackupFileSize    = 50
)

func main() {

	confighelper.InitViper()

	maxBackupCount, _ = strconv.Atoi(confighelper.GetConfig("maxBackupCount"))
	maxBackupFileSize, _ = strconv.Atoi(confighelper.GetConfig("maxBackupFileSize"))
	maxAgeForBackupFiles, _ = strconv.Atoi(confighelper.GetConfig("maxAgeForBackupFiles"))

	logginghelper.Init("D:/logs/server.log", false, maxBackupCount, maxBackupFileSize, maxAgeForBackupFiles, true)

	e := echo.New()
	api.Init(e)
	apiServerPort = confighelper.GetConfig("apicdnServerPort")
	e.Start(":" + apiServerPort)
}
