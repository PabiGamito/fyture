package auth

import (
	"fyture/controllers"
	"fyture/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetAuthToken(c *gin.Context) {
	session := sessions.Default(c)
	UID := session.Get("UID").(uint)

	db := c.MustGet("db").(*gorm.DB)
	user, err := controllers.FindUserByID(UID, db)
	if err != nil {
		c.JSON(403, "User not found.")
		return
	}

	token, err := createAccessToken(&user)

	c.JSON(200, token)
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
