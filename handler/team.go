package handler

import (
	"log"
	"net/http"
	"scrim-api/model"
	"scrim-api/service"

	"github.com/gin-gonic/gin"
)

func HandlerTeamCreate(c *gin.Context, userId string) {
	var data model.TeamCreateReq

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	teamId, err := service.TeamCreate(userId, data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, teamId)
}

func HandlerTeamUpdate(c *gin.Context) {
	var data model.TeamUpdate

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := service.TeamUpdate(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func HandlerTeamJoin(c *gin.Context) {
	var data model.TeamJoin

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	resp, err := service.TeamJoin(data)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func HandlerTeamMemberGet(c *gin.Context, teamId string) {
	resp, err := service.TeamMemberGet(teamId)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func HandlerTeamDetailGet(c *gin.Context, teamId string) {
	resp, err := service.TeamDetailGet(teamId)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func HandlerTeamInviteCodeGet(c *gin.Context, teamId string) {
	resp, err := service.TeamSetFlagInviteCode(teamId)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func HandlerTeamDelete(c *gin.Context, teamId string) {
	err := service.TeamDelete(teamId)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"error_code": "0000",
		"error_msg":  "successfully delete team",
	})
}
