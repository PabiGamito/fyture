package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"fyture/controllers"
	"fyture/models"
)

type GoogleProfile struct {
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
}

func GoogleLogin(c *gin.Context) {
	url := location.Get(c)
	googleOAuthConfig := googleOAuthConfig(url.Scheme, url.Host)

	// Redirect user to Google's consent page.
	redirectURL := googleOAuthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
}

func GoogleLoginCallback(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")
	url := location.Get(c)
	googleOAuthConfig := googleOAuthConfig(url.Scheme, url.Host)
	userProfile, err := getUserProfile(googleOAuthConfig, code)

	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/auth/google")
		return
	}

	var user models.User
	db := c.MustGet("db").(*gorm.DB)

	// Look for user in database
	user, err = controllers.FindUserByEmail(userProfile.Email, db)

	if err != nil {
		// Create new user
		input := &controllers.CreateUserInput{
			FirstName:    userProfile.GivenName,
			LastName:     userProfile.FamilyName,
			Email:        userProfile.Email,
			ProfileImage: userProfile.Picture,
		}

		user = controllers.CreateUser(input, db)
	}

	err = setUserSession(c, &user)
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/auth/google")
	}

	c.JSON(200, userProfile)
}

func googleOAuthConfig(scheme string, host string) oauth2.Config {
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).
	return oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  fmt.Sprintf("%s://%s/auth/callback/google", scheme, host),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func getUserProfile(googleOAuthConfig oauth2.Config, code string) (GoogleProfile, error) {
	var profile GoogleProfile

	token, err := googleOAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return profile, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return profile, fmt.Errorf("Failed getting Google user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return profile, fmt.Errorf("Failed reading Google OAuth user info response body: %s", err.Error())
	}

	if err := json.Unmarshal(contents, &profile); err != nil {
		return profile, fmt.Errorf("Failed parsing Google OAuth user info response JSON: %s", err.Error())
	}

	return profile, nil
}
