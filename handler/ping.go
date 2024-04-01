package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @summary ping
// @produce plain
// @Tags         ping
// @response 200 {string} string "pong"
// @router /ping [get]
func HandlerPing(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
