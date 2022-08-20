package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func User(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	r.Static("/user/image", "./Images")
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
				"success": true,
				"message": "password " + user.Username + " berhasil diperbarui",
				"error":   nil,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Password yang baru keduanya tidak sama.",
			})
			return
		}
	})

	r.PATCH("/profile-picture", Auth.Authorization(), func(c *gin.Context) {
		id, _ := c.Get("id")

		var student Model.Student
		var newAvatar Model.Student

		avatar, _ := c.FormFile("avatar")

		if avatar != nil {
			rand.Seed(time.Now().Unix())

			str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

			shuff := []rune(str)

			rand.Shuffle(len(shuff), func(i, j int) {
				shuff[i], shuff[j] = shuff[j], shuff[i]
			})
			avatar.Filename = string(shuff)

			if err := c.SaveUploadedFile(avatar, "./Images/"+avatar.Filename); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Success": false,
					"error":   "upload file err: " + err.Error(),
				})
				return
			}

			godotenv.Load("../.env")
			newAvatar = Model.Student{
				Avatar:      os.Getenv("BASE_URL") + "/api/user/image/" + avatar.Filename,
			}
		} else {
			if res := db.Where("id = ?", id).Take(&student); res.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Error when updating the database.",
					"error":   res.Error.Error(),
				})
			}
			newAvatar = Model.Student{
				Avatar:      student.Avatar,
			}
		}

		result := db.Where("id = ?", id).Model(&newAvatar).Updates(newAvatar)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when updating the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		if result = db.Where("id = ?", id).Take(&newAvatar); result.Error != nil {
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
				"message": "group not found.",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Update successful.",
			"data":    newAvatar.Avatar,
		})

	})
}
