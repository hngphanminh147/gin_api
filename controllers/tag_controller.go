package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hngphanminh147/gin_api/db"
	"github.com/hngphanminh147/gin_api/models"
)

type CreateTagInput struct {
	Name string `gorm:"type:nvarchar(255)" json:"name" binding:"required"`
	// Post []*models.Post `gorm:"many2many:post_tag" json:"post" binding:"required"`
}

type UpdateTagInput struct {
	Name string `gorm:"type:nvarchar(255)" json:"name" binding:"required"`
	// Post []*models.Post `gorm:"many2many:post_tag" json:"post"`
}

// GET /tags
func FindTags(c *gin.Context) {
	var tags []models.Tag

	db.Db.Find(&tags)

	c.JSON(http.StatusOK, gin.H{"data": tags})
}

// POST /tags
func CreateTag(c *gin.Context) {
	var input CreateTagInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Tag{
		Name: input.Name,
		// Post: input.Post,
	}
	db.Db.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// UPDATE /tag/:id
func UpdateTag(c *gin.Context) {
	var tag models.Tag
	if err := db.Db.Where("id = ?", c.Param("id")).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
	}

	var input UpdateTagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	db.Db.Model(&tag).Updates(models.Tag{Name: input.Name})
	// db.Db.Model(&tag).Association("Post").Replace(input.Post)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}
