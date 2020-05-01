package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var featureRequests = []string{"Feature 1", "Feature 2"}

	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"featureRequests": featureRequests,
	})
}
