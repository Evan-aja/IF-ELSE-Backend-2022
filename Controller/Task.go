package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Task(db *gorm.DB, q *gin.Engine) {
	// Biar lebih gampang
	r := q.Group("/api")

	// GetTaskById
	r.GET("task/:AgendaId", func(c *gin.Context) {
		AgendaId, isIdExists := c.Params.Get("AgendaId")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "id is not available",
			})
			return
		}

		var displayData Model.Task

		err := db.Where("agenda_id", AgendaId).First(&displayData).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error reading database",
				"success": false,
				"Error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Title":       displayData.Title,
			"Description": displayData.Description,
			"Agenda Id":   displayData.AgendaId,
			"Condition":   displayData.Condition,
			"Step":        displayData.Step,
		})
	})

	// Edit task
	r.PATCH("/task/edit/:AgendaId", Auth.Authorization(), func(c *gin.Context) {
		task := Model.Task{}
		AgendaId, isIdExists := c.Params.Get("AgendaId")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "id is not available",
			})
			return
		}

		input := Model.Task{}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "cannot get input",
				"error":   err.Error(),
			})
			return
		}

		in := Model.Task{
			Title:       input.Title,
			Description: input.Description,
			AgendaId:    input.AgendaId,
			Condition:   input.Condition,
			Step:        input.Step,
		}

		result := db.Where("agenda_id = ?", AgendaId).Model(&in).Updates(in).Error
		if result != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when updating the database.",
				"error":   result.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":     true,
			"error":       nil,
			"message":     "success",
			"title":       task.Title,
			"Description": task.Description,
			"AgendaId":    task.AgendaId,
			"Condition":   task.Condition,
			"Step":        task.Step,
		})
	})

	// Submit task
	//r.PATCH("/task/:AgendaId", Auth.Authorization(), submitTask)
	r.PATCH("/task/:AgendaId", Auth.Authorization(), func(c *gin.Context) {
		task := Model.Task{}
		AgendaId, isIdExists := c.Params.Get("AgendaId")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "id is not available",
			})
			return
		}

		input := Model.Task{}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "cannot get input",
				"error":   err.Error(),
			})
			return
		}

		in := Model.Task{
			Fields:      input.Fields,
			IsPublished: true,
		}

		result := db.Where("agenda_id = ?", AgendaId).Model(&in).Updates(in).Error
		if result != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when updating the database.",
				"error":   result.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":     true,
			"error":       nil,
			"message":     "success",
			"title":       task.Title,
			"Description": task.Description,
			"AgendaId":    task.AgendaId,
			"Condition":   task.Condition,
			"Step":        task.Step,
			"StartAt":     task.StartAt,
			"EndAt":       task.EndAt,
		})
	})

	// Membuat task baru
	r.POST("/task", func(c *gin.Context) {
		var input Model.Task
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "cannot get input",
				"error":   err.Error(),
			})
			return
		}

		task := Model.Task{
			Title:       input.Title,
			Description: input.Description,
			AgendaId:    input.AgendaId,
			Condition:   input.Condition,
			Step:        input.Step,
			StartAt:     input.StartAt,
			EndAt:       input.EndAt,
		}

		if input.Title == "" || input.Description == "" || input.AgendaId == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "evevrything must be filled",
			})
			return
		}

		if err := db.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error on creating task",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success":     true,
			"message":     "Task created successfully",
			"Title":       task.Title,
			"Description": task.Description,
			"AgendaId":    task.AgendaId,
			"Condition":   task.Condition,
			"Step":        task.Step,
		})
	})

	// Menunjukan semua task
	//r.GET("/task/all", getAllTask)
	r.GET("/task/all", func(c *gin.Context) {
		var tasks []Model.Task
		if err := db.Find(&tasks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"error":   nil,
			"message": "success",
			"data": gin.H{
				"tasks": tasks,
			},
		})
	})

	// Delete Task
	//r.POST("/task/deleteTask", deleteTask)
	// Delete Task (cara lama)
	r.POST("/task/delete", func(c *gin.Context) {
		var input Model.Task

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "cannot get input",
				"error":   err.Error(),
			})
			return
		}

		del := Model.Task{
			AgendaId: input.AgendaId,
		}

		if err := db.Where("agenda_id = ?", del.AgendaId).First(&del).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"succsess": false,
				"message":  "error finding data",
				"error":    err.Error(),
			})
			return
		}

		deleted := db.Delete(&del).Error
		if deleted != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"succsess": false,
				"message":  "error deleting data",
				"error":    deleted.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "table is deleted",
		})
	})
}

func TaskUser(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/Task")

	r.POST("/T", func(c *gin.Context) {
		
	})

	r.POST("/", func(c *gin.Context) {
		var tasks []Model.Task

		if err := db.Find(&tasks).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    tasks,
		})
	})
}

// var db *gorm.DB

// func deleteTask(c *gin.Context) {
// 	var input Model.Task
// 	AgendaId, isIdExists := c.Params.Get("AgendaId")
// 	if !isIdExists {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"Success": false,
// 			"message": "id is not available",
// 		})
// 		return
// 	}

// 	if err := c.BindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": "cannot get input",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	del := Model.Task{
// 		AgendaId: input.AgendaId,
// 	}

// 	if err := db.Where("agenda_id = ?", AgendaId).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"succsess": false,
// 			"message":  "error accessing database",
// 			"error":    err.Error(),
// 		})
// 		return
// 	}

// 	deleted := db.Delete(&del).Error
// 	if deleted != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"succsess": false,
// 			"message":  "error deleting data",
// 			"error":    deleted.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "table is deleted",
// 	})
// }

// func addTask(c *gin.Context) {
// 	var input Model.Task
// 	if err := c.BindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": "cannot get input",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	task := Model.Task{
// 		Title:       input.Title,
// 		Description: input.Description,
// 		AgendaId:    input.AgendaId,
// 		StartAt:     input.StartAt,
// 		EndAt:       input.EndAt,
// 	}

// 	if input.Title == "" || input.Description == "" || input.AgendaId == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": "evevrything must be filled",
// 		})
// 		return
// 	}

// 	if err := db.Create(&task).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"success": false,
// 			"message": "error on creating task",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success":     true,
// 		"message":     "Task created successfully",
// 		"Title":       task.Title,
// 		"Description": task.Description,
// 		"AgendaId":    task.AgendaId,
// 	})
// }

// func submitTask(c *gin.Context) {
// 	task := Model.Task{}
// 	AgendaId, isIdExists := c.Params.Get("AgendaId")
// 	if !isIdExists {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"Success": false,
// 			"message": "id is not available",
// 		})
// 		return
// 	}

// 	input := Model.Task{}

// 	if err := c.BindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"success": false,
// 			"message": "cannot get input",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	in := Model.Task{
// 		Fields:      input.Fields,
// 		IsPublished: true,
// 	}

// 	result := db.Where("agenda_id = ?", AgendaId).Model(&in).Updates(in).Error
// 	if result != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"success": false,
// 			"message": "Error when updating the database.",
// 			"error":   result.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success":     true,
// 		"error":       nil,
// 		"message":     "success",
// 		"title":       task.Title,
// 		"Description": task.Description,
// 		"AgendaId":    task.AgendaId,
// 		"StartAt":     task.StartAt,
// 		"EndAt":       task.EndAt,
// 	})
// }

// func getAllTask(c *gin.Context) {
// 	var tasks []Model.Task
// 	if err := db.Find(&tasks).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"success": false,
// 			"message": "Something went wrong",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"error":   nil,
// 		"message": "success",
// 		"data": gin.H{
// 			"tasks": tasks,
// 		},
// 	})
// }

// Mencari task berdasarkan id
// r.GET("/task/:AgendaId", func(c *gin.Context) {
// 	task := Model.Task{}
// 	AgendaId, isIdExists := c.Params.Get("AgendaId")
// 	if !isIdExists {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"Success": false,
// 			"message": "id is not available",
// 		})
// 		return
// 	}

// 	if err := db.Where("agenda_id = ?", AgendaId).Take(&task).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"success": false,
// 			"message": "Error reading the database.",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"success":     true,
// 		"error":       nil,
// 		"message":     "success",
// 		"title":       task.Title,
// 		"Description": task.Description,
// 		"AgendaId":    task.AgendaId,
// 		"StartAt":     task.StartAt,
// 		"EndAt":       task.EndAt,
// 	})
// })
