package handler

import (
	"log"
	"net/http"
	"scrim-api/service"

	"github.com/gin-gonic/gin"
)

func HandlerMapNameGet(c *gin.Context, gameId string) {
	resp, err := service.MapNameGet(gameId)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
