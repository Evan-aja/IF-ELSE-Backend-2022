package Controller

import (
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Agenda(db *gorm.DB, q *gin.Engine) {

	r := q.Group("/api")
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

        c.JSON(http.StatusOK, gin.H {
            "success": true,
            "message": "query completed.",
            "data": agenda,
        })
	})

}
