package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ifelse/Auth"
	"ifelse/Model"
	"net/http"
	"reflect"
)

func PenugasanUser(db *gorm.DB, q *gin.Engine) {
	r := q.Group("user/penugasan")

	r.GET("/ppp", func(c *gin.Context) {
		var penugasan []Model.Penugasan

		if err := db.Find(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    penugasan,
			"success": true,
		})
	})

	r.GET("/", Auth.Authorization(), func(c *gin.Context) {
		//id, _ := c.MustGet("id").(uint)

		id, _ := c.MustGet("id").(uint)

		//fmt.Println(id)
		//fmt.Println()
		//fmt.Println(reflect.TypeOf(id))

		//coba pake model.student

		var input Model.Penugasan

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "cannot get input",
			})
			return
		}

		// float64

		//ms := Model.Links{
		//	Link1:     input.Link1,
		//	Link2:     input.Link2,
		//	StudentID: id,
		//}

		fmt.Println(reflect.TypeOf(id))

		//if err := db.Where("id = ?", id).Model(&ms).Updates(ms).Error; err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{
		//		"message": err.Error(),
		//		"success": false,
		//		"msg":     "cannot update data",
		//	})
		//	return
		//}

		c.JSON(http.StatusOK, gin.H{
			"ID": id,
		})
	})
}
