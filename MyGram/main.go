package main

import (
	"mygram/config"
	"mygram/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	//users

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	//socialMedias
	r.POST("/socialMedias", controllers.CreateSocialMedia)
	r.GET("/socialMedias", controllers.GetSocialMedias)
	r.PUT("/socialMedias/:id", controllers.UpdateSocialMedia)
	r.DELETE("/socialMedia/:id", controllers.DeleteSocialMedia)

	//photos
	r.POST("/photos", controllers.CreatePhoto)
	r.GET("/photos", controllers.GetPhotos)
	r.PUT("/photos/:id", controllers.UpdatePhoto)
	r.DELETE("/photos/:id", controllers.DeletePhoto)

	//comments
	r.POST("/comment", controllers.CreateComment)
	r.GET("/comment", controllers.GetComments)
	r.PUT("/comment/:id", controllers.UpdateComment)
	r.DELETE("/comment/:id", controllers.DeleteComment)

	r.Run(":8080")
}
