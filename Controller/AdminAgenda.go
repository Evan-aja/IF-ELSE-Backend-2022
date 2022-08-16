package Controller

import (
	"ifelse/Model"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func AdminAgenda(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	r.Static("/agenda/image", "./Images")
	r.POST("/admin/agenda", func(c *gin.Context) {
		image, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "get form err: " + err.Error(),
			})
			return
		}

		rand.Seed(time.Now().Unix())

		str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		shuff := []rune(str)

		rand.Shuffle(len(shuff), func(i, j int) {
			shuff[i], shuff[j] = shuff[j], shuff[i]
		})
		image.Filename = string(shuff)

		godotenv.Load("../.env")

		if err := c.SaveUploadedFile(image, "./Images/"+image.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"error":   "upload file err: " + err.Error(),
			})
			return
		}

		newAgenda := Model.Agenda{
			Title:            c.PostForm("title"),
			Content:          c.PostForm("content"),
			Image:            os.Getenv("BASE_URL") + "/api/agenda/image/" + image.Filename,
			StartAt:          c.PostForm("start_at"),
			EndAt:            c.PostForm("end_at"),
			PerizinanStartAt: c.PostForm("perizinan_start_at"),
			PerizinanEndAt:   c.PostForm("perizinan_end_at"),
			IsPublished:      false,
		}

		if err := db.Create(&newAgenda); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when inserting a new agenda",
				"error":   err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success":     true,
			"message":     "a new agenda has successfully created",
			"error":       nil,
			"nama_agenda": newAgenda.Title,
		})
	})

	r.GET("/admin/agenda", func(c *gin.Context) {
		var agenda []Model.Agenda
		if result := db.Find(&agenda); result.Error != nil {
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
			"data":    agenda,
		})
	})
}
