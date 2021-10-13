package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hngphanminh147/gin_api/db"
	"github.com/hngphanminh147/gin_api/models"
)

type CreatePostInput struct {
	Title string        `gorm:"type:nvarchar(255)" json:"title" binding:"required"`
	Body  string        `gorm:"type:text" json:"body" binding:"required"`
	Tag   []*models.Tag `gorm:"many2many:post_tag" json:"tag" binding:"required"`
}

type UpdatePostInput struct {
	Title string        `gorm:"type:nvarchar(255)" json:"title"`
	Body  string        `gorm:"type:text" json:"body"`
	Tag   []*models.Tag `gorm:"many2many:post_tag" json:"tag"`
}

// GET /posts
func FindPosts(c *gin.Context) {
	var posts []models.Post

	db.Db.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// POST /posts
func CreatePost(c *gin.Context) {
	var input CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:      input.Title,
		Body:       input.Body,
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Tag:        input.Tag,
	}
	db.Db.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// POST /posts/:id
func FindPost(c *gin.Context) {
	var post models.Post
	if err := db.Db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// PUT /post/:id
func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := db.Db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Db.Model(&post).Updates(models.Post{
		Title:      input.Title,
		Body:       input.Body,
		Updated_at: time.Now(),
	})
	db.Db.Model(&post).Association("Tag").Replace(input.Tag)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DELETE /post/:id
func DeletePost(c *gin.Context) {
	var post models.Post
	if err := db.Db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Db.Model(&post).Association("Tag").Clear()
	db.Db.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}
