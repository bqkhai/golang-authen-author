package controllers

import (
	"authen-author-example/common"
	"authen-author-example/models"
	"authen-author-example/utils"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Signup(ctx *gin.Context) {
	var signUpInput models.SignUpInput

	if err := ctx.BindJSON(&signUpInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// compare password
	if signUpInput.Password != signUpInput.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error in confirm password",
		})
		return
	}

	// hash password
	pass, _ := utils.HashPassword(signUpInput.Password)

	// check email exist
	var count int64
	common.Instance.Model(&models.User{}).Where("email = ?", signUpInput.Email).Count(&count)
	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is already exist",
		})
		return
	}

	// create user
	newUser := models.User{
		Email:    signUpInput.Email,
		Password: pass,
	}
	result := common.Instance.Create(&newUser)

	// response
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := models.UserResponse{
		Id:        newUser.Id,
		FullName:  newUser.FullName,
		Email:     newUser.Email,
		Role:      newUser.Role,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, response)
}

func Login(ctx *gin.Context) {
	var signInInput models.SignInInput
	if err := ctx.BindJSON(&signInInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check email
	var user models.User
	result := common.Instance.Model(&models.User{}).Where("email = ?", signInInput.Email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Email incorrect"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Check if the password is correct
	if error := utils.VerifyPassword(user.Password, signInInput.Password); error != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Email or password incorrect"})
		return
	}

	// Generate the JWT token string
	ttl, err := time.ParseDuration(os.Getenv("JWT_TOKEN_EXPIRED"))
	if err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	tokenString, err := utils.GenerateToken(ttl, signInInput.Email, os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"access_token": tokenString})
}
