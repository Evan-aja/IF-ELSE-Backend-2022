package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminTask(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	// Get All Penugasan
	r.GET("/admin/task", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId > 1 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		}
		var task []Model.Task

		if err := db.Preload("Links").Find(&task).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}

		if task == nil {
			c.JSON(http.StatusOK, gin.H{
				"Error":   "task isn't available",
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
	r.GET("/admin/task/:id", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId > 1 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		}
		
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

	// Create Penugasan
	r.POST("/admin/task", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId > 1 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		} 

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
				studentTask.LinkPos = int32(j)
				studentTask.LinkID = linkId[j]
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
	r.PATCH("/admin/task/:id", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId > 1 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		} 
		id := c.Param("id")

		var ntask Model.NewTask
		var task Model.Task
		var oldTask Model.Task
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

		if err := db.Where("id = ?", id).Take(&oldTask); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't find task",
				"success": false,
				"error":   err.Error.Error(),
			})
			return
		}

		task = Model.Task{
			ID: oldTask.ID,
			Title:       ntask.Title,
			Description: ntask.Description,
			Condition:   ntask.Condition,
			Step:        ntask.Step,
			JumlahLink:  ntask.JumlahLink,
			Deadline:    ntask.Deadline,
		}

		if err := db.Where("id = ? ", id).Model(&task).Select("*").Updates(task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't updated task",
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		link = Model.Links{
			TaskID: task.ID,
		}

		var allLink []Model.Links

		if res := db.Where("task_id = ?", id).Delete(&allLink); res.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "can't delete old links",
				"success": false,
				"error":   res.Error.Error(),
			})
			return
		}

		if ntask.JumlahLink > oldTask.JumlahLink {
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
					studentTask.LinkPos = int32(j)
					studentTask.LinkID = linkId[j]
					studentTask.ID = 0
					if err := db.Create(&studentTask).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"message": "can't create studentTask",
							"success": false,
							"error":   err.Error(),
						})
						return
					}
				}
			}
		} else if ntask.JumlahLink < oldTask.JumlahLink {
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
			for i := 0; i < len(student); i++ {
				for j := 0; j < int(task.JumlahLink); j++ {
					studentTask.StudentID = student[i].ID
					studentTask.LinkPos = int32(j)
					studentTask.LinkID = linkId[j]
					studentTask.ID = 0
					if err := db.Create(&studentTask).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"message": "can't create studentTask",
							"success": false,
							"error":   err.Error(),
						})
						return
					}
				}
			}
		} else {
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
					studentTask.LinkPos = int32(j)
					studentTask.LinkID = linkId[j]
					studentTask.ID = 0
					if err := db.Create(&studentTask).Error; err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"message": "can't create studentTask",
							"success": false,
							"error":   err.Error(),
						})
						return
					}
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    ntask,
			"message": "successfully updated a task data",
			"success": true,
		})
	})

	// Delete Penugasan by Id
	r.DELETE("/admin/task/:id", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId > 1 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		} 
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is not supplied.",
			})
			return
		}
		parsedId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is invalid.",
			})
			return
		}
		task := Model.Task{
			ID: uint(parsedId),
		}
		if result := db.Delete(&task); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when deleting from the database.",
				"error":   result.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Delete successful.",
		})
	})
}
