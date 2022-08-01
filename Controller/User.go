package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	// show logged in user profile
	r.GET("/profile", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.Get("id")
		user := Model.User{}
		if err := db.Where("id=?", id).Preload("Student").Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"error":   nil,
			"message": "success",
			"data": gin.H{
				"id":       user.ID,
				"name":     user.Name,
				"username": user.Username,
				"email":    user.Email,
				"student":  user.Student,
			},
		})
	})

}
