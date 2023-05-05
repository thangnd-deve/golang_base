package core

import (
	"github.com/gin-gonic/gin"
	"thesis/app/database"
	"thesis/app/helpers"
)

var (
	router = gin.Default()
)

func StartApp() {
	gin.SetMode(helpers.Env("GIN_MODE", "debug"))
	database.Migration()
	MapUrls()
	_ = router.Run(":" + helpers.Env("APP_PORT", "8080"))
}
