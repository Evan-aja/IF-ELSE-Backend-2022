package Controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ifelse/Model"
	"net/http"
)

func Admin(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	// untuk menampilkan seluruh data mahasiswa yang tersedia
	// ditambah fitur search dengan menggunakan nama atau nim
	r.POST("/mahasiswa", func(c *gin.Context) {
		name, _ := c.GetQuery("name")
		nim, _ := c.GetQuery("nim")

		var queryResults []Model.Student

		if res := db.Where("name LIKE ?", "%"+name+"%").Where("nim LIKE ?", "%"+nim).Find(&queryResults); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Query is not supplied.",
				"error":   res.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Search successful",
			"data": gin.H{
				"query": gin.H{
					"name": name,
					"nim":  nim,
				},
				"result": queryResults,
			},
		})

	})

	// untuk menampilkan data mahasiswa berdasarkan id yang diminta
	r.GET("/mahasiswa/:id", func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "user_id is not available",
			})
			return
		}

		mahasiswa := Model.Student{}

		if result := db.Where("id = ?", id).Take(&mahasiswa); result.Error != nil {
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
			"data":    mahasiswa,
		})

	})

	// untuk menampilkan data mahasiswa yang memiliki `group_id` yang sama
	r.GET("/mahasiswa/group/:id", func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "group_id is not available",
			})
			return
		}

		var mahasiswa []Model.Student

		if result := db.Where("group_id = ?", id).Find(&mahasiswa); result.Error != nil {
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
			"data":    mahasiswa,
		})

	})

	// untuk memperbarui data `group_id` mahasiswa
	r.PATCH("/mahasiswa/group/:id", func(c *gin.Context) {
		id, _ := c.Params.Get("id")

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
			GroupId: body.GroupId,
		}
		result := db.Where("id = ?", id).Updates(mahasiswa)
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
