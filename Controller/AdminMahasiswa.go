package Controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ifelse/Model"
	"net/http"
	"strconv"
)

func AdminMahasiswa(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	// untuk menampilkan seluruh data mahasiswa yang tersedia
	// ditambah fitur search dengan menggunakan nama atau nim
	r.POST("/admin/mahasiswa", func(c *gin.Context) {
		name, _ := c.GetQuery("name")
		nim, _ := c.GetQuery("nim")

		q := c.Request.URL.Query()

		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 1:
			pageSize = 10
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize

		var queryResults []Model.Student

		if res := db.Where("nickname LIKE ?", "%"+name+"%").Where("nim LIKE ?", "%"+nim+"%").Offset(offset).Limit(pageSize).Find(&queryResults); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "students is not found.",
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
	r.GET("/admin/mahasiswa/:id", func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "id is not available",
			})
			return
		}

		var mahasiswa Model.Student

		if result := db.Where("id = ?", id).Take(&mahasiswa); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		stask := []Model.StudentTask{}
		// link := []Model.Links{}

		// preload task, mahasiswa, dan links
		if result := db.Where("student_id = ?", id).Preload("Task").Preload("Student").Find(&stask); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		smark := []Model.Marking{}

		if result := db.Where("student_id = ?", id).Find(&smark); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		// var link Model.Links 

		// if res := db.Where("task_id = ?", id).Find(&link); res.Error != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"success": false,
		// 		"message": "Error when querying the database.",
		// 		"error":   res.Error.Error(),
		// 	})
		// 	return
		// }


		
		// if result := db.Where("task_id = ?", stask.ID).Find(&link); result.Error != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"success": false,
		// 		"message": "Error when querying the database.",
		// 		"error":   result.Error.Error(),
		// 	})
		// 	return
		// }
		// data := Model.Group{}

		// if group := db.Where("id = ?", mahasiswa.GroupID).Preload("Student").Take(&data); group.Error != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"success": false,
		// 		"message": "Error when querying the database.",
		// 		"error":   group.Error.Error(),
		// 	})
		// 	return
		// }

		// var mark []Model.StudentMarking

		// if rang := db.Where("student_id = ?", mahasiswa.ID).Find(&mark); rang.Error != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"success": false,
		// 		"message": "Error when querying the database.",
		// 		"error":   rang.Error.Error(),
		// 	})
		// 	return
		// }
		// stask.Links = link[0]
		mahasiswa.StudentTask = stask
		mahasiswa.Marking = smark
		// mahasiswa.StudentMarking = mark

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    mahasiswa,
		})

	})

	// untuk memperbarui data `group_id` mahasiswa
	r.PATCH("/admin/mahasiswa/:id", func(c *gin.Context) {
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
			GroupID: body.GroupID,
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

	// Menghapus mahasiswa berdasrkan ID yang dimiliki
	r.DELETE("/mahasiswa/:id", func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is not supplied.",
			})
			return
		}
		parsedId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is invalid.",
			})
			return
		}
		student := Model.Student{
			ID: uint(parsedId),
		}
		if result := db.Delete(&student); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when deleting from the database.",
				"error":   result.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Delete successful.",
		})
	})
}
