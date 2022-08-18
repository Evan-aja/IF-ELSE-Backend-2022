package Controller

import (
	"fmt"
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminTask(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	// Get All Penugasan
	r.GET("/admin/task", func(c *gin.Context) {
		var task []Model.Task

		if err := db.Find(&task).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}

		if task == nil {
			c.JSON(http.StatusOK, gin.H{
				"Error":   "Data Is Empty",
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    task,
			"success": true,
		})
	})

	// Get Penugasan By ID
	r.GET("/admin/task/:id", func(c *gin.Context) {
		var penugasan Model.Task

		id := c.Param("id")

		if err := db.Where("id = ?", id).First(&penugasan).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"message": "Something went wrong on server side",
				"success": false,
			})
			return
		}

		if penugasan.Title == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": penugasan,
				"error":   "Data Is Empty",
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": penugasan,
			"success": true,
			"error":   nil,
		})
	})

	// Create Penugasan
	r.POST("/admin/new-task", func(c *gin.Context) {
		var ntask Model.NewTask
		var task Model.Task
		var link Model.Links

		if err := c.BindJSON(&ntask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
				"message": "Can't read input",
			})
			return
		}

		if ntask.Title == "" || ntask.Description == "" || ntask.Condition == "" || ntask.Step == "" || ntask.Links == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "All field bust be filled",
				"success": false,
				"error":   nil,
			})
			return
		}
		var student []Model.Student
		if err := db.Find(&student); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error when querying database",
				"success": false,
				"error":   err.Error.Error(),
			})
		}

		task = Model.Task{
			Title:       ntask.Title,
			Description: ntask.Description,
			Condition:   ntask.Condition,
			Step:        ntask.Step,
			JumlahLink:  ntask.JumlahLink,
			Deadline:    ntask.Deadline,
		}
		// buat task baru
		if err := db.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't created task",
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		link = Model.Links{
			TaskID: task.ID,
		}
		// buat link yang di dalam task
		var linkId []uint
		for i := 0; i < len(ntask.Links); i++ {
			link.Title = ntask.Links[i]
			link.ID = 0
			if err := db.Create(&link); err.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "can't create links",
					"success": false,
					"error":   err.Error.Error(),
				})
				return
			}
			linkId = append(linkId, link.ID)
		}

		studentTask := Model.StudentTask{
			TaskID: task.ID,
		}
		// assign link ke siswa
		for i := 0; i < len(student); i++ {
			for j := 0; j < int(task.JumlahLink); j++ {
				studentTask.StudentID = student[i].ID
				studentTask.LinkID = linkId[j]
				fmt.Println(studentTask)
				studentTask.ID = 0
				if err := db.Create(&studentTask).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "can't create links",
						"success": false,
						"error":   err.Error(),
					})
					return
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    ntask,
			"success": true,
		})
	})

	// Update Penugasan
	r.PATCH("/edit/:id", func(c *gin.Context) {
		var penugasan Model.Task

		id := c.Param("id")

		if err := db.Where("id = ?", id).First(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't create data",
			})
			return
		}

		if err := c.BindJSON(&penugasan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't read input",
			})
			return
		}

		if err := db.Save(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't update data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    penugasan,
			"success": true,
		})
	})

	// Delete Penugasan by Id
	r.DELETE("/delete", func(c *gin.Context) {
		var penugasan Model.Task

		if err := c.BindJSON(&penugasan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't read input",
			})
			return
		}

		id := penugasan.ID

		if err := db.Where("id = ?", id).First(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't find data",
			})
			return
		}

		if err := db.Delete(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't delete data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    penugasan,
			"success": true,
		})
	})

	// Delete All Penugasan
	//r.DELETE("/delete/all", func(c *gin.Context) {
	//	var penugasan []Model.Penugasan
	//
	//	if err := db.Find(&penugasan).Error; err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": err.Error(),
	//			"success": false,
	//			"Reason":  "Can't find data",
	//		})
	//		return
	//	}
	//
	//	if err := db.Delete(&penugasan).Error; err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": err.Error(),
	//			"success": false,
	//			"Reason":  "Can't delete data",
	//		})
	//		return
	//	}

	//	c.JSON(http.StatusOK, gin.H{
	//		"data":    penugasan,
	//		"success": true,
	//	})
	//})

	// Delete All Penugasan
	//r.DELETE("/delete/all/:id", func(c *gin.Context) {
	//	var penugasan []Model.Penugasan
	//
	//	id := c.Param("id")
	//
	//	if err := db.Where("id = ?", id).Find(&penugasan).Error; err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": err.Error(),
	//			"success": false,
	//			"Reason":  "Can't find data",
	//		})
	//		return
	//	}
	//
	//	if err := db.Delete(&penugasan).Error; err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": err.Error(),
	//			"success": false,
	//			"Reason":  "Can't delete data",
	//		})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"data":    penugasan,
	//		"success": true,
	//	})
	//})
}
