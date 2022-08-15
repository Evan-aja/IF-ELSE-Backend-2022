package Controller

import (
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BeritaUser(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/user/berita")

	r.GET("/", func(c *gin.Context) {
		var berita []Model.Berita

		if err := db.Find(&berita).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "Data Is Empty",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"berita": berita,
		})
	})
}
