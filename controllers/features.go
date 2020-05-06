package controllers

import (
	"fmt"
	"fyture/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateFeatureInput struct {
	Title   string `json:"title" binding:"required"`
	Details string `json:"details" binding:"required"`
}

type UpdateFeatureInput struct {
	Title   string `json:"title"`
	Details string `json:"details"`
}

// FindFeatures Get all features
func FindFeatures(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	params := c.Request.URL.Query()

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil || limit > 100 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing limit parameter"})
		return
	}

	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing offset parameter"})
		return
	}

	var features []models.Feature
	switch params.Get("sort") {
	case "trending":
		db.Limit(limit).Offset(offset).Find(&features)
	case "top":
		db.Order("up_votes desc").Limit(limit).Offset(offset).Find(&features)
	case "new":
		db.Order("created_at desc").Limit(limit).Offset(offset).Find(&features)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing sort parameter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"features": features})
}

func FindFeature(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var feature models.Feature
	if err := db.Where("id = ?", c.Param("id")).First(&feature).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"feature": feature})
}

func CreateFeature(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input CreateFeatureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Feature
	feature := models.Feature{
		Title:     input.Title,
		Details:   input.Details,
		UpVotes:   1,
		CreatedAt: time.Now().Unix(),
	}
	db.Create(&feature)

	c.JSON(http.StatusOK, gin.H{"id": feature.ID})
}

func UpdateFeature(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var feature models.Feature
	if err := db.Where("id = ?", c.Param("id")).First(&feature).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateFeatureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&feature).Updates(input)

	c.Status(http.StatusOK)
}

func DeleteFeature(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var feature models.Feature
	if err := db.Where("id = ?", c.Param("id")).First(&feature).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&feature)

	c.Status(http.StatusOK)
}

func UpVote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	fmt.Println("Ran 1")

	err := db.Transaction(func(tx *gorm.DB) error {
		fmt.Println("Ran 2")

		// Get model if exist
		var feature models.Feature
		if err := tx.Where("id = ?", c.Param("id")).First(&feature).Error; err != nil {
			return err
		}

		fmt.Println("Ran 3")

		// Update vote count
		if err := tx.Model(&feature).Update("UpVotes", feature.UpVotes+1).Error; err != nil {
			return err
		}

		fmt.Println("Ran 4")

		return nil
	})

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
