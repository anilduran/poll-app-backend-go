package routes

import (
	"example.com/poll-app-backend-go/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	auth := r.Group("/api/auth")
	auth.POST("/sign-in", SignIn)
	auth.POST("/sign-up", SignUp)

	user := r.Group("/api/users")
	user.Use(middlewares.Auth)
	user.GET("/", GetUsers)
	user.GET("/:id", GetUserByID)
	user.POST("/", CreateUser)
	user.PUT("/:id", UpdateUser)
	user.DELETE("/:id", DeleteUser)

	poll := r.Group("/api/polls")
	poll.Use(middlewares.Auth)
	poll.GET("/", GetPolls)
	poll.GET("/:id", GetPollByID)
	poll.POST("/", CreatePoll)
	poll.PUT("/:id", UpdatePoll)
	poll.DELETE("/:id", DeletePoll)

	poll.POST("/:id/vote", VotePoll)
	poll.DELETE("/:id/vote", UnvotePoll)

	option := r.Group("/api/options")
	option.Use(middlewares.Auth)
	option.GET("/options", GetOptions)
	option.POST("/options", CreateOption)
	option.PUT("/options/:id", UpdateOption)
	option.DELETE("/options/:id", DeleteOption)

	categories := r.Group("/api/categories")
	categories.Use(middlewares.Auth)
	categories.GET("/", GetCategories)
	categories.GET("/:id", GetCategoryByID)
	categories.POST("/", CreateCategory)
	categories.PUT("/:id", UpdateCategory)
	categories.DELETE("/:id", DeleteCategory)

}
