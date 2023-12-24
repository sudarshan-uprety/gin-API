package controller

import (
	"net/http"
	"os"
	"time"

	"API/database"
	"API/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user = &models.User{}

		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to hash password",
			})
			return
		}

		user.Password = string(hash)

		result := database.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error,
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "User registration successful", "email": user.Email,
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body struct {
			Email    string
			Password string
		}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var user = &models.User{}

		database.DB.First(&user, "email = ?", body.Email)
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			})
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid password",
			})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to create a token",
			})
			return
		}
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})

	}

}

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")

		userData := struct {
			Name  string `json:"name"`
			Age   int    `json:"age"`
			Phone string `json:"phone"`
			Email string `json:"email"`
		}{
			Name:  user.(models.User).Name,
			Age:   user.(models.User).Age,
			Phone: user.(models.User).Phone,
			Email: user.(models.User).Email,
		}

		c.JSON(http.StatusOK, gin.H{
			"data": userData,
		})
	}
}
