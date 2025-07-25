package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/abrishk26/a2sv-project-track/task7/Domain"
	"github.com/abrishk26/a2sv-project-track/task7/Infrastructure"
	"github.com/abrishk26/a2sv-project-track/task7/Repositories"
	"github.com/abrishk26/a2sv-project-track/task7/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserController(userColl *mongo.Collection) *UserController {
	return &UserController{userColl}
}

type UserController struct {
	userColl *mongo.Collection
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

	userUsecase := usecases.NewUserUsecases(repositories.NewUserRepository(c.Request.Context(), uc.userColl))

	user, err := userUsecase.Get(idParam)
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

	userUsecase := usecases.NewUserUsecases(repositories.NewUserRepository(c.Request.Context(), uc.userColl))
	err = userUsecase.Add(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser, err := userUsecase.Get(user.ID)
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
	userUsecase := usecases.NewUserUsecases(repositories.NewUserRepository(c.Request.Context(), uc.userColl))
	users, err := userUsecase.GetAll()
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

	userUsecase := usecases.NewUserUsecases(repositories.NewUserRepository(c.Request.Context(), uc.userColl))
	user, err := userUsecase.Get(idParam)
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

	userUsecase := usecases.NewUserUsecases(repositories.NewUserRepository(c.Request.Context(), uc.userColl))
	err = userUsecase.Update(idParam, user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	updatedUser, err := userUsecase.Get(idParam)
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

	userUsecase := usecases.NewUserUsecases(repositories.NewUserRepository(c.Request.Context(), uc.userColl))
	deletedUser, err := userUsecase.Get(idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = userUsecase.Delete(idParam)
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
