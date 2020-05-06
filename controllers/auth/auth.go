package auth

import (
	"fmt"
	"fyture/controllers"
	"fyture/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AuthenticationRequired(c *gin.Context) {
	// TODO: Try adding caching this (redis) since this middleware might be called a lot
	user, err := findUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Valid authentication required"})
		c.Abort()
		return
	}

	c.Set("user", &user)

	c.Next()
}

func GetAuthToken(c *gin.Context) {
	user, err := findUserFromSession(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Missing session or user not found.")
		c.Abort()
		return
	}

	token, err := createAccessToken(&user)

	c.JSON(200, token)
}

func findUserFromSession(c *gin.Context) (models.User, error) {
	session := sessions.Default(c)
	UID := session.Get("UID")

	if UID == nil {
		return models.User{}, fmt.Errorf("missing UID session")
	}

	db := c.MustGet("db").(*gorm.DB)
	return controllers.FindUserByID(UID.(uint), db)
}

func setUserSession(c *gin.Context, user *models.User) error {
	session := sessions.Default(c)

	session.Set("UID", user.ID)
	err := session.Save()

	return err
}

func createAccessToken(user *models.User) (string, error) {
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("JTW_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}
