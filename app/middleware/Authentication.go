package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"thesis/app/helpers"
	"time"
)

type Middleware struct {
}

func (middleware Middleware) Auth(ctx *gin.Context) {

	httpClient := &http.Client{Timeout: time.Duration(1) * time.Second}

	headerToken := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")
	accessToken := strings.Join(headerToken, " ")
	accessToken = strings.TrimSpace(accessToken)

	url := helpers.Env("AUTH_HOST", "localhost:8000") + "/verify-token"
	req, err := http.NewRequest("POST", url, nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)
	response, err := httpClient.Do(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		ctx.Abort()
		return
	}
	if response.StatusCode != 200 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		ctx.Abort()
		return
	}
	ctx.Next()
}
