package Controller

import (
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminMarking(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	r.PATCH("/admin/marking/:id", func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		var body Model.Marking

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "body is invalid",
				"error":   err.Error(),
			})
			return
		}

		// ID, _ := strconv.ParseUint(id, 10, 64)

		marking := Model.Marking{
			Mark: body.Mark,
		}
		result := db.Where("student_id = ?", id).Where("agenda_id = ?", body.AgendaID).Updates(marking)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when updating the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		var mahasiswa Model.Student

		if result = db.Where("id = ?", id).Preload("Marking").Take(&mahasiswa); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		if result.RowsAffected < 1 {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Mahasiswa not found.",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Update successful.",
			"data":    mahasiswa,
		})
	})
}
