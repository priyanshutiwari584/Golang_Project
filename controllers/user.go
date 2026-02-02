package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"go-project.com/go-project/models"
	"go-project.com/go-project/utils"
)

func Singup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		utils.BadRequest(context, "Could not parse request body", err)
		return
	}

	err = user.Save()

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			utils.Conflict(context, "User already exists!")
			return
		}

		utils.InternalServerError(context, err, "Could not save user")
		return
	}

	utils.Created(context, "User created successfully!", nil)
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		utils.BadRequest(context, "Could not parse request body", err)
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		utils.Unauthorized(context, "Invalid credentials")
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		utils.InternalServerError(context, err, "Could not authenticate user")
	}

	utils.OK(context, "User logged in successfully!", gin.H{"token": token})
}
