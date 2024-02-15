package controllers

import (
	"net/http"
	"todo_project_be/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create task",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"data":    task,
	})
}
