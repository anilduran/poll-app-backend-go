package routes

import (
	"net/http"

	"example.com/poll-app-backend-go/db"
	"example.com/poll-app-backend-go/models"
	"github.com/gin-gonic/gin"
)

func GetPolls(c *gin.Context) {

	var polls []models.Poll

	result := db.DB.Find(&polls)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": polls,
	})

}

func GetPollByID(c *gin.Context) {

	var poll models.Poll

	id := c.Param("id")

	result := db.DB.First(&poll, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, poll)

}

func CreatePoll(c *gin.Context) {

	type CreatePollInput struct {
		Question string `form:"question" binding:"required"`
		UserID   uint   `form:"user_id" binding:"required"`
	}

	var input CreatePollInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	poll := models.Poll{Question: input.Question, UserID: input.UserID}

	result := db.DB.Create(&poll)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, poll)

}

func UpdatePoll(c *gin.Context) {

	type UpdatePollInput struct {
		Question string `form:"question" binding:"required"`
		UserID   uint   `form:"user_id" binding:"required"`
	}

	var input UpdatePollInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	var poll models.Poll

	id := c.Param("id")

	result := db.DB.First(&poll, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	poll.Question = input.Question
	poll.UserID = input.UserID

	result = db.DB.Save(&poll)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, poll)

}

func DeletePoll(c *gin.Context) {

	var poll models.Poll

	id := c.Param("id")

	result := db.DB.First(&poll, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	result = db.DB.Delete(&poll, id)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, poll)

}

func VotePoll(c *gin.Context) {

}

func UnvotePoll(c *gin.Context) {

}
