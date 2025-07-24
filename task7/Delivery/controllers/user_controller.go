package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"
	"github.com/abrishk26/a2sv-project-track/task7/Infrastructure"
	"github.com/abrishk26/a2sv-project-track/task7/Repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewUserController(um *repositories.UserRepository) *UserController {
	return &UserController{um}
}

type UserController struct {
	userRepository *repositories.UserRepository
}

func (uc *UserController) LoginUser(c *gin.Context) {
	idParam := c.Param("id")

	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user id missing",
		})
		return
	}

	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	user, err := uc.userRepository.Get(c, idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = infrastructures.CompareHashAndPassword([]byte(user.PasswordHash), []byte(requestBody.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	claims := map[string]any{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add((24 * time.Hour)).Unix(),
	}

	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "signing key not found",
		})
		return
	}

	tokenString, err := infrastructures.GenerateJWT(claims, secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unexpected error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "logged in successfully",
		"token_string": tokenString,
	})
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var requestBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	hashedPassword, err := infrastructures.HashPassword(requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	userID, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user domain.User
	user.ID = userID.String()
	user.PasswordHash = string(hashedPassword)
	user.Username = requestBody.Username

	err = uc.userRepository.Add(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser, err := uc.userRepository.Get(c, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"user":    newUser,
	})
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userRepository.GetAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "users retrieved successfully",
		"tasks":   users,
	})
}

func (uc *UserController) GetUser(c *gin.Context) {
	idParam := c.Param("id")

	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user id missing",
		})
		return
	}

	user, err := uc.userRepository.Get(c, idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user retrieved successfully",
		"task":    user,
	})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id missing",
		})
		return
	}

	var user domain.User

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err = uc.userRepository.Update(c, idParam, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedUser, err := uc.userRepository.Get(c, idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "user updated successfully",
		"task":    updatedUser,
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id missing",
		})
		return
	}

	deletedUser, err := uc.userRepository.Get(c, idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = uc.userRepository.Delete(c, idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "user deleted successfully",
		"task":    deletedUser,
	})
}
