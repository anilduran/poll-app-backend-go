package routes

import (
	"net/http"

	"example.com/poll-app-backend-go/db"
	"example.com/poll-app-backend-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	userId := c.GetUint("userId")

	if userId == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	type CreatePollInput struct {
		Question string    `form:"question" binding:"required"`
		UserID   uuid.UUID `form:"user_id" binding:"required"`
	}

	var input CreatePollInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	poll := models.Poll{Question: input.Question, UserID: userId}

	result := db.DB.Create(&poll)

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, poll)

}

func UpdatePoll(c *gin.Context) {

	type UpdatePollInput struct {
		Question string    `form:"question" binding:"required"`
		UserID   uuid.UUID `form:"user_id" binding:"required"`
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

	userId := c.GetUint("userId")

	if userId == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	type VotePollInput struct {
		OptionID uint `form:"option_id" binding:"required"`
	}

	var input VotePollInput

	err := c.ShouldBind(&input)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	vote := models.Vote{UserID: userId, OptionID: input.OptionID}

	result := db.DB.Create(&vote)

	if result != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, vote)

}

func UnvotePoll(c *gin.Context) {
	userId := c.GetUint("userId")

	if userId == 0 {
		c.Status(http.StatusUnauthorized)
		return
	}

	pollId := c.Param("id")

	result := db.DB.Where("user_id = ? AND poll_id = ?", userId, pollId).Delete(&models.Vote{})

	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(200)
}
