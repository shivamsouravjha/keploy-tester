package controllers

import (
	"fmt"
	"net/http"
	"segwise/clients/postgres"
	"segwise/config"
	"segwise/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Struct for request payloads
type AuthInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Secret key for JWT
var jwtSecret = []byte(config.Get().JWT_SECRET)

// @Summary Register a new user
// @Description Creates a new user with a username and password
// @Accept  json
// @Produce  json
// @Param request body AuthInput true "User Registration Data"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
// @Security BearerAuth
func RegisterUser(c *gin.Context) {
	db := postgres.GetDB()

	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{Username: input.Username, Password: string(hashedPassword)}
	db.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})

}

// @Summary User login
// @Description Authenticates a user and returns a JWT token
// @Accept  json
// @Produce  json
// @Param request body AuthInput true "User Login Credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Security BearerAuth
// @Router /login [post]
func LoginUser(c *gin.Context) {
	db := postgres.GetDB()

	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		fmt.Println("Error: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	fmt.Println([]byte(user.Password), []byte(input.Password), "asd")
	// Verify password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := generateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

// Generate JWT Token
func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
