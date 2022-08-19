package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserPerizinan(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	r.POST("/perizinan", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.MustGet("id").(uint)
		
		// studentId := uint(id)
		var body Model.Perizinan

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "body is invalid",
				"error":   err.Error(),
			})
			return
		}

		newPerizinan := Model.Perizinan {
			StudentID: id,
			AgendaID: body.AgendaID,
			LinkSurat: body.LinkSurat,
		}

		if err := db.Create(&newPerizinan); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when inserting a new permission link",
				"error":   err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "a new group has successfully created",
			"error":   nil,
			"judul":   newPerizinan,
		})


	})
}
