package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminMahasiswa(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	type Student struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		GroupName string `json:"group_name"`
		NIM       string `json:"nim"`
	}
	// untuk menampilkan seluruh data mahasiswa yang tersedia
	// ditambah fitur search dengan menggunakan nama atau nim
	r.POST("/admin/mahasiswa", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId == 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}
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

		var lengthStudents []Model.Student

		if res := db.Find(&lengthStudents); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "students is not found.",
				"error":   res.Error.Error(),
			})
			return
		}

		var queryResults []Model.Student

		if res := db.Where("name LIKE ?", "%"+name+"%").Where("nim LIKE ?", "%"+nim+"%").Where("name NOT LIKE ?", "%admin%").Offset(offset).Limit(pageSize).Find(&queryResults); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "students is not found.",
				"error":   res.Error.Error(),
			})
			return
		}

		var group []Model.Group
		var ret []Student

		for i := 0; i < len(queryResults); i++ {
			if result := db.Where("id = ?", queryResults[i].GroupID).Find(&group); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Error when querying the database.",
					"error":   result.Error.Error(),
				})
				return
			}
			var temp Student
			temp.ID = queryResults[i].ID
			temp.Name = queryResults[i].Name
			temp.GroupName = group[0].GroupName
			temp.NIM = queryResults[i].NIM
			ret = append(ret, temp)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    ret,
			"length":  len(lengthStudents),
		})

		// for  i := 0, element := index quequeryResults {

		// }

		// c.JSON(http.StatusOK, gin.H{
		// 	"success": true,
		// 	"message": "Search successful",
		// 	"data": gin.H{
		// 		"query": gin.H{
		// 			"name": name,
		// 			"nim":  nim,
		// 		},
		// 		"result": queryResults,
		// 	},
		// })

	})

	// untuk menampilkan data mahasiswa berdasarkan id yang diminta
	r.GET("/admin/mahasiswa/:id", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId == 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}
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
		var group Model.Group

		if result := db.Where("id = ?", mahasiswa.GroupID).Find(&group); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		var stask []Model.StudentTask

		// preload task, mahasiswa, dan links
		if result := db.Where("student_id = ?", id).Find(&stask); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		var task []Model.Task
		type StudentTask struct {
			ID          uint      `json:"id"`
			TaskID      uint      `json:"task_id"`
			TaskTitle   string    `json:"task_title"`
			LabelLink   string    `json:"label_link"`
			Link        string    `json:"link"`
			SubmittedAt time.Time `json:"time"`
		}

		var ret []StudentTask
		var temp StudentTask

		for i := 0; i < len(stask); i++ {
			if res := db.Where("id = ?", &stask[i].TaskID).Preload("Links").Find(&task); res.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "can't found task",
					"success": false,
					"error":   res.Error.Error(),
				})
				return
			}
			temp.ID = stask[i].ID
			temp.Link = stask[i].Link
			temp.LabelLink = task[0].Links[stask[i].LinkPos].Title
			temp.TaskID = task[0].ID
			temp.TaskTitle = task[0].Title
			temp.SubmittedAt = stask[i].SubmittedAt
			ret = append(ret, temp)
		}
		smark := []Model.Marking{}

		if result := db.Where("student_id = ?", id).Preload("Agenda").Find(&smark); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		mahasiswa.GroupName = group.GroupName
		mahasiswa.Marking = smark

		c.JSON(http.StatusOK, gin.H{
			"success":      true,
			"message":      "query completed.",
			"data":         mahasiswa,
			"student_task": ret,
		})

	})

	// untuk memperbarui data `group_id` mahasiswa
	r.PATCH("/admin/mahasiswa/:id", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId != 0 && user.RoleId != 3 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}
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

	// Menghapus mahasiswa berdasarkan ID yang dimiliki
	r.DELETE("admin/mahasiswa/:id", Auth.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user Model.User
		if err := db.Where("id = ?", ID).Take(&user); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error.Error(),
			})
			return
		}

		if user.RoleId != 0 && user.RoleId != 3 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}

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
