package handler

import (
	"log"
	"net/http"
	"scrim-api/model"
	"scrim-api/service"

	"github.com/gin-gonic/gin"
)

func HandlerRegister(c *gin.Context) {
	var userRegInfo model.UserRegisterReq

	if err := c.ShouldBindJSON(&userRegInfo); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := service.UserRegister(userRegInfo)

	if err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func HandlerLogin(c *gin.Context) {
	var userLoginData model.UserLoginReq

	if err := c.ShouldBindJSON(&userLoginData); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	resp, err := service.UserLogin(userLoginData)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "0001",
			"error_msg":  "not found user, maybe email and password are incorrect.  ",
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func HandlerChangeRole(c *gin.Context) {
	var data model.ChangeRole

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.Status(http.StatusBadRequest)
		return
	}

	err := service.ChangeRole(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "0001",
			"error_msg":  "error changing role.",
		})
		return
	}

	c.Status(http.StatusOK)
}

func HandlerKickMember(c *gin.Context) {
	var data model.KickMember

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_msg": err.Error(),
		})
		return
	}

	err := service.KickMember(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "0001",
			"error_msg":  "error kicking member.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": "0000",
		"error_msg":  "successfully kick.",
	})
}

func HandlerUpdateUserProfile(c *gin.Context, userId string) {
	var data model.UserUpdateData

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_msg": err.Error(),
		})
		return
	}

	err := service.UpdateUserProfile(userId, data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "0001",
			"error_msg":  "error updating user profile.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": "0000",
		"error_msg":  "successfully updated.",
	})
}

func HandlerForgotPassword(c *gin.Context) {
	var data model.ForgotPasswordReq

	if err := c.ShouldBindJSON(&data); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_msg": err.Error(),
		})
		return
	}

	err := service.SendPasswordToEmail(data)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error_code": "0001",
			"error_msg":  "error sending email or email does not exist in the system.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error_code": "0000",
		"error_msg":  "successfully sent password to email.",
	})
}
