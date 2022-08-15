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

	r.PATCH("/profile", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.Get("id")

		var body Model.Student

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "body is invalid",
				"error":   err.Error(),
			})
			return
		}

		mahasiswa := Model.Student{
			Nickname: body.Nickname,
			Address:  body.Address,
			Line:     body.Line,
			Whatsapp: body.Whatsapp,
			About:    body.About,
		}

		result := db.Where("id = ?", id).Model(&mahasiswa).Updates(mahasiswa)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when updating the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		if result = db.Where("id = ?", id).Take(&mahasiswa); result.Error != nil {
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
				"message": "mahasiswa not found.",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "successfully updated data.",
			"data":    mahasiswa,
		})
	})

	r.PATCH("/password", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.Get("id")

		var body Model.ChangePassword

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "body is invalid",
				"error":   err.Error(),
			})
			return
		}

		user := Model.User{}

		compare := body.Newpass1 == body.Newpass2

		if compare {
			user = Model.User{
				Password: hash(body.Newpass1),
				ID: user.ID,
			}

			result := db.Where("id = ? ", id).Model(&user).Updates(&user)

			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Error when updating the database.",
					"error":   result.Error.Error(),
				})
				return
			}

			if result = db.Where("id = ? ", id).Take(&user); result.Error != nil {
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
					"message": "user not found.",
				})
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"success":  true,
				"message":  "password berhasil diperbarui",
				"data":     user.Name,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Password yang baru keduanya tidak sama.",
			})
			return
		}
	})

}
