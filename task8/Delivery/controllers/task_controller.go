package controllers

import (
	"net/http"

	"github.com/abrishk26/a2sv-project-track/task8/Domain"
	"github.com/abrishk26/a2sv-project-track/task8/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewTaskController(usecase *usecases.TaskUsecases) *TaskController {
	return &TaskController{usecase}
}

type TaskController struct {
	usecase *usecases.TaskUsecases
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task domain.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	taskID, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	task.ID = taskID.String()

	err = tc.usecase.Add(c.Request.Context(), task)
	if err != nil {
		switch err {
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		case domain.ErrDuplicateTask:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "task already exists",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"task": task,
	})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.usecase.GetAll(c.Request.Context())
	if err != nil {
		switch err {
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
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
		"message": "tasks retrieved successfully",
		"tasks":   tasks,
	})
}

func (tc *TaskController) GetTask(c *gin.Context) {
	idParam := c.Param("id")

	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id missing",
		})
		return
	}

	task, err := tc.usecase.Get(c.Request.Context(), idParam)
	if err != nil {
		switch err {
		case domain.ErrTaskNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorize",
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
		"message": "task retrieved successfully",
		"task":    task,
	})
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id missing",
		})
		return
	}

	var task domain.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err = tc.usecase.Update(c.Request.Context(), idParam, task)
	if err != nil {
		switch err {
		case domain.ErrTaskNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorize",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	updatedTask, err := tc.usecase.Get(c.Request.Context(), idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "task updated successfully",
		"task":    updatedTask,
	})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "task id missing",
		})
		return
	}

	err := tc.usecase.Delete(c.Request.Context(), idParam)
	if err != nil {
		switch err {
		case domain.ErrTaskNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		case domain.ErrUnauthorized:
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorize",
			})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "task deleted successfully",
	})
}
