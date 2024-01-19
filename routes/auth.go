package routes

import (
	"net/http"

	"example.com/poll-app-backend-go/db"
	"example.com/poll-app-backend-go/models"
	"example.com/poll-app-backend-go/utils"
	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {

	type SignInInput struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var input SignInInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	isPasswordTrue := utils.ComparePassword(input.Password, user.Password)

	if !isPasswordTrue {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong password",
		})
		return
	}

	var token string

	token, err = utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func SignUp(c *gin.Context) {

	type SignUpInput struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	var input SignUpInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var existingUser models.User

	result := db.DB.Where("email = ?", input.Email).First(&existingUser)

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})
		return
	}

	hashedPassword := utils.HashPassword(input.Password)

	user := models.User{Email: input.Email, Password: hashedPassword}

	result = db.DB.Create(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var token string

	token, err = utils.GenerateToken(user.ID, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't generate token",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})

}
