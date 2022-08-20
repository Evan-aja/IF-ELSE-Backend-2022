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
			"data":    task,
			"success": true,
			"error":   nil,
		})
	})

	r.PATCH("/task", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.MustGet("id").(uint)

		var body Model.NewTask

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "body is invalid",
				"error":   err.Error(),
			})
			return
		}
		var slink []Model.StudentTask
		var task Model.Task

		for i := 0; i < len(body.Links); i++ {
			if res := db.Where("task_id = ?", body.ID).Find(&slink); res.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "can't create links",
					"success": false,
					"error":   res.Error.Error(),
				})
				return
			}
			slink[i].Link = body.Links[i]
			slink[i].StudentID = id


			if res := db.Updates(&slink[i]); res.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "can't updates links",
					"success": false,
					"error":   res.Error.Error(),
				})
				return
			}

			if res := db.Where("id = ?", slink[i].TaskID).Take(&task); res.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "can't found task",
					"success": false,
					"error":   res.Error.Error(),
				})
				return
			}
		}

		type StudentTask struct {
			ID uint `json:"id"`
			TaskTitle string `json:"task_title"`
			Link string `json:"link"`
			UpdatedAt time.Time `json:"time"`
		}

		var ret[] StudentTask
		for _, element := range slink {
			var temp StudentTask
			temp.ID = element.ID
			temp.Link = element.Link
			temp.TaskTitle = task.Title
			temp.UpdatedAt = element.UpdatedAt
			ret = append(ret, temp)
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "a new link has successfully uploaded",
			"error":   nil,
			"data":    ret,
		})
	})
}
