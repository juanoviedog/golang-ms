package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/juanoviedog/golang-ms/initializers"
	"github.com/juanoviedog/golang-ms/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	// Get the data from the req body
	var body struct {
		Email    string
		password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.password), 10)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
	}

	// Respond
	ctx.JSON(http.StatusOK, gin.H{"message": "User Created"})
}

func Login(ctx *gin.Context) {
	// Get the data from the req body
	var body struct {
		Email    string
		password string
	}

	if ctx.Bind(&body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user or password"})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user or password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign it
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token"})
		return
	}

	// Send it back
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{})

}

func Validate(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"message": user})
}
