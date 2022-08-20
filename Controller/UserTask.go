package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserTask(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	r.GET("/task", func(c *gin.Context) {
		var task []Model.Task

		if result := db.Preload("Links").Find(&task); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    task,
		})
	})

	r.GET("/task/:id", func(c *gin.Context) {
		var task Model.Task

		id := c.Param("id")

		if err := db.Where("id = ?", id).Preload("Links").First(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Something went wrong on server side",
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": task,
			"success": true,
			"error":   nil,
		})
	})


	r.PATCH("/task", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.MustGet("id").(uint)

		var body Model.StudentTask

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "body is invalid",
				"error":   err.Error(),
			})
			return
		}

		signTask := Model.StudentTask {
			ID: body.ID,
			StudentID: id,
			Link: body.Link,
			UpdatedAt: time.Now(),
		} 
		if err := db.Updates(&signTask); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when inserting a new permission link",
				"error":   err.Error.Error(),
			})
			return
		}

		if err := db.Preload("Task").Preload("Links.Task").Preload("Student").Take(&signTask); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when query data",
				"error":   err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "a new group has successfully created",
			"error":   nil,
			"data":   signTask,
		})
	})
}
