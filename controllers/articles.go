// controllers/articles

package controllers

import (
	"github.com/gerrr/blog_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /articles
// Get all articles
func AllArticles(c *gin.Context) {
	var article []models.Article
	models.DB.Find(&article)
	
	c.JSON(http.StatusOK, gin.H{"data": article})
}


// POST /articles
type CreateArticleInput struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content string `json:"content" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// Create new article
func CreateArticle(c *gin.Context) {
	// Validate input
	var input CreateArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	article := models.Article{
		Title: input.Title,
		Description: input.Description,
		Content: input.Content,
		Author: input.Author,
	}
	models.DB.Create(&article)
	
	c.JSON(http.StatusOK, gin.H{"data": article})
}


// GET /article/:id
// Get some article
func FindArticle(c *gin.Context) {
	var article models.Article
	
	if err := models.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Article not found!"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": article})
}


// DELETE /article/:id
// Delete the article
func DeleteArticle(c *gin.Context) {
	var article models.Article
	
	if err := models.DB.Where("id = ?", c.Param("id")).First(&article).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Article not found!"})
		return
	}
	
	models.DB.Delete(&article)
	c.JSON(http.StatusOK, gin.H{"data": true})
}