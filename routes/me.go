package routes

import (
	"net/http"

	"example.com/poll-app-backend-go/db"
	"example.com/poll-app-backend-go/models"
	"example.com/poll-app-backend-go/utils"
	"github.com/gin-gonic/gin"
)

func GetMyCredentials(c *gin.Context) {

	userId := c.GetUint("userId")

	if userId == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	var user models.User

	result := db.DB.First(&user, userId)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func UpdateMyCredentials(c *gin.Context) {

	userId := c.GetUint("userId")

	if userId == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	type UpdateMyCredentialsInput struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}

	var input UpdateMyCredentialsInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var user models.User

	result := db.DB.Where("email = ?", input.Email).First(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	user.Email = input.Email

	hashedPassword := utils.HashPassword(input.Password)

	user.Password = hashedPassword

	result = db.DB.Save(&user)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}

func GetMyPolls(c *gin.Context) {

	userID := c.GetUint("userId")

	if userID == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	var polls []models.Poll

	result := db.DB.Where("user_id = ?", userID).Find(&polls)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, polls)

}
