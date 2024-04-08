package handler

import (
	"log"
	"net/http"
	"scrim-api/model"
	"scrim-api/service"

	"github.com/gin-gonic/gin"
)

// @summary HandlerScrimPost
// @Param request body  true model.ScrimPost "request body"
// @Success 201 {int}
// @Tags         team
// @response 201 {int}
// @router /team [post]
func HandlerScrimPost(c *gin.Context) {
	var data model.ScrimPost

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	resp, err := service.ScrimPost(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func HandlerScrimMakeOffer(c *gin.Context) {
	var data model.ScrimMakeOffer

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := service.ScrimMakeOffer(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": "0000",
		"error_msg":  "successfully make offer",
	})
}

func HandlerScrimAcceptOffer(c *gin.Context) {
	var data model.ScrimAcceptOffer

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := service.ScrimAcceptOffer(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": "0000",
		"error_msg":  "successfully accepted",
	})
}

func HandlerScrimCancelMatch(c *gin.Context) {
	var data model.ScrimCancelMatch

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := service.ScrimCancelMatch(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": "0000",
		"error_msg":  "successfully cancel",
	})
}

func HandlerScrimDelete(c *gin.Context) {
	var data model.ScrimDelete

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := service.ScrimDelete(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"error_code": "0000",
		"error_msg":  "successfully delete scrim",
	})
}

func HandlerScrimGetOffer(c *gin.Context, teamId string) {
	resp, err := service.ScrimGetOffer(teamId)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func HandlerScrimQuery(c *gin.Context) {
	var data model.ScrimGetReq

	if err := c.BindQuery(&data); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := service.ScrimGet(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func HandlerScrimGetMatch(c *gin.Context, teamId string) {
	resp, err := service.ScrimGetMatch(teamId)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
