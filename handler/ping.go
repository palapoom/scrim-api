package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerPing(c *gin.Context) {
	c.Status(http.StatusOK)
}
