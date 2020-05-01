package main

import (
	"feature-requests/controllers"
	"feature-requests/models"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Server Vue app
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	// r.NoRoute(func(c *gin.Context) {
	// 	c.File("./dist/index.html")
	// })

	// Feature API
	api := r.Group("/api")
	{
		api.GET("/features", controllers.FindFeatures)
		api.GET("/feature/:id", controllers.FindFeature)
		api.POST("/feature", controllers.CreateFeature)
		api.PATCH("/feature/:id", controllers.UpdateFeature)
		api.DELETE("/feature/:id", controllers.DeleteFeature)

		api.POST("/feature/upvote/:id", controllers.UpVote)
	}

	r.Run("0.0.0.0:8080")
}
