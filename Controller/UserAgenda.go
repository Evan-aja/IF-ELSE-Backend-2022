package Controller

import (
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserAgenda(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	type Agenda struct {
		ID      uint   `json:"id"`
		Title   string `json:"title"`
		Image   string `json:"image"`
		StartAt string `json:"start_at"`
		EndAt   string `json:"end_at"`
	}

	r.GET("/agenda", func(c *gin.Context) {
		var agenda []Model.Agenda
		if result := db.Find(&agenda); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}
		var ret []Agenda

		for _, value := range agenda {
			var temp Agenda
			temp.ID = value.ID
			temp.Title = value.Title
			temp.Image = value.Image
			temp.StartAt = value.StartAt
			temp.EndAt = value.EndAt
			ret = append(ret, temp)
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    ret,
		})
	})
}
