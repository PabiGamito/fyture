package main

import (
	"os"

	"github.com/gin-contrib/location"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"fyture/controllers"
	"fyture/controllers/auth"
	"fyture/models"
)

func main() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")

	// configure to automatically detect scheme and host
	// - use http when default scheme cannot be determined
	// - use localhost:8080 when default host cannot be determined
	r.Use(location.Default())

	// Setup sessions
	store := cookie.NewStore([]byte(os.Getenv("COOKIE_SECRET")))
	r.Use(sessions.Sessions("session", store))

	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Server Vue app
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	r.GET("/feature/:id", func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	// Auth
	r.GET("/auth/google", auth.GoogleLogin)
	r.GET("/auth/callback/google", auth.GoogleLoginCallback)

	r.GET("/auth/getToken", auth.GetAuthToken)

	// API
	api := r.Group("/api")
	// Unauthenticated endpoints
	{
		api.GET("/features", controllers.FindFeatures)
		api.GET("/feature/:id", controllers.FindFeature)
	}
	// Authenticated endpoints
	{
		authenticatedAPI := api.Group("/")
		authenticatedAPI.Use(auth.AuthenticationRequired)
		{
			authenticatedAPI.POST("/feature", controllers.CreateFeature)
			authenticatedAPI.PATCH("/feature/:id", controllers.UpdateFeature)
			authenticatedAPI.DELETE("/feature/:id", controllers.DeleteFeature)

			authenticatedAPI.POST("/feature/upvote/:id", controllers.UpVote)
		}

	}

	r.Run("0.0.0.0:8080")
}
