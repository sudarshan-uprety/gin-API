package controller

import (
	"API/database"
	"API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var posts []models.Post
		if err := database.DB.Preload("User").Where("visiblity = ?", true).Find(&posts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})

	}
}

func GetPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		var posts []models.Post
		userIdStr := c.Param("id")
		userID, err := strconv.ParseUint(userIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if err := database.DB.Preload("User").Where("visiblity = ? AND user_id = ?", true, userID).Find(&posts).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{
			"data": posts,
		})
	}
}

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post = &models.Post{}
		// var user = &models.User{}
		if err := c.Bind(&post); err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}
		post.UserID = user.(models.User).ID
		result := database.DB.Create(&post)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Post created successfully",
		})

	}
}

func UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post = &models.Post{}
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		postIdStr := c.Param("id")
		postID, err := strconv.ParseUint(postIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := database.DB.Where("id = ?", postID).First(&post).Error; err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}
		if user.(models.User).ID != post.UserID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "This post doesnot belongs to you.",
			})
			return
		}

		if err := database.DB.Save(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Post updated successfully.",
			"data":    post,
		})
	}
}

func DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post = &models.Post{}
		postIdStr := c.Param("id")
		PostID, err := strconv.ParseUint(postIdStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}
		if err := database.DB.Where("id = ?", PostID).Find(&post).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if user.(models.User).ID != post.UserID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "This post doesnot belongs to you.",
			})
			return
		}
		if err := database.DB.Delete(&post).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Post deleted successfully.",
		})
	}
}
