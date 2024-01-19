package routes

import (
	"net/http"

	"example.com/poll-app-backend-go/db"
	"example.com/poll-app-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetOptions(c *gin.Context) {

	var options []models.Option

	result := db.DB.Find(&options)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": options,
	})

}

func CreateOption(c *gin.Context) {

	type CreateOptionInput struct {
		Text   string    `form:"text" binding:"required"`
		PollID uuid.UUID `form:"poll_id" binding:"required"`
	}

	var input CreateOptionInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	option := models.Option{Text: input.Text, PollID: input.PollID}

	result := db.DB.Create(&option)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, option)

}

func UpdateOption(c *gin.Context) {

	type UpdateOptionInput struct {
		Text string `form:"text"`
	}

	var input UpdateOptionInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var option models.Option

	id := c.Param("id")

	result := db.DB.First(&option, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	if input.Text != "" {
		option.Text = input.Text
	}

	result = db.DB.Save(&option)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, option)

}

func DeleteOption(c *gin.Context) {

	var option models.Option

	id := c.Param("id")

	result := db.DB.First(&option, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&option)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, option)

}
