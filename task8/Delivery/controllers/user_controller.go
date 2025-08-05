package controllers

import (
	"net/http"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
	"github.com/abrishk26/a2sv-project-track/task8/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewUserController(usecase *usecases.UserUsecases) *UserController {
	return &UserController{usecase}
}

type UserController struct {
	usecase *usecases.UserUsecases
}

func (uc *UserController) LoginUser(c *gin.Context) {
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

	token, err := uc.usecase.Login(c.Request.Context(), requestBody.Username, requestBody.Username)
	if err != nil {
		switch err {
		case domain.ErrInvalidCredential:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid username or password",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "logged in successfully",
		"token_string": token,
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

	userID, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	var user domain.User
	user.ID = userID.String()
	user.PasswordHash = requestBody.Password
	user.Username = requestBody.Username

	err = uc.usecase.Register(c.Request.Context(), user)
	if err != nil {
		switch err {
		case domain.ErrDuplicateUsername:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "username already exists",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	newUser, err := uc.usecase.Get(c.Request.Context(), user.ID)
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
	users, err := uc.usecase.GetAll(c.Request.Context())
	if err != nil {
		switch err {
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unautorized",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
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

	user, err := uc.usecase.Get(c.Request.Context(), idParam)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
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

	err = uc.usecase.Update(c.Request.Context(), idParam, user)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	updatedUser, err := uc.usecase.Get(c.Request.Context(), idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
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

	err := uc.usecase.Delete(c.Request.Context(), idParam)
	if err != nil {
		switch err {
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "user deleted successfully",
	})
}
