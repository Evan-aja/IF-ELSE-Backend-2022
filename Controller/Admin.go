package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Admin(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/admin")

	// Admin Kelompok 
	// Bagian Data Mahasiswa
	// untuk menampilkan seluruh data mahasiswa yang tersedia (search using query name or nim and pagination)
	r.POST("/mahasiswa", Auth.Authorization(), func(c *gin.Context) {
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

		if res := db.Where("name LIKE ?", "%"+name+"%").Where("nim LIKE ?", "%"+nim).Offset(offset).Limit(pageSize).Find(&queryResults); res.Error != nil {
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
	r.GET("/mahasiswa/:id", Auth.Authorization(), func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "id is not available",
			})
			return
		}

		mahasiswa := Model.Student{}

		if result := db.Where("id = ?", id).Preload("StudentTask").Take(&mahasiswa); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		data := Model.Group{}
		
		if group := db.Where("id = ?", mahasiswa.GroupID).Take(&data); group.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   group.Error.Error(),
			})
			return
		}

		mark := []Model.StudentMarking{}

		if rang := db.Where("student_id = ?", mahasiswa.ID).Find(&mark); rang.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   rang.Error.Error(),
			})
			return
		}

		mahasiswa.Group = data
		mahasiswa.StudentMarking = mark

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    mahasiswa,
		})

	})

	// untuk memperbarui data `group_id` mahasiswa
	r.PATCH("/mahasiswa/:id", Auth.Authorization(), func(c *gin.Context) {
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
	r.DELETE("/mahasiswa/:id", Auth.Authorization(), func(c *gin.Context) {
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

	// Group Routers
	r.Static("/assets", "./assets")

	type Group struct {
		ID uint `json:"id"`
		CompanionName string `json:"companion_name"`
		GroupName string `json:"group_name"`
	}
	// untuk menampilkan seluruh grup yang ada
	r.GET("/group", func(c *gin.Context) {
		group := []Model.Group{}

		if result := db.Preload("Student").Find(&group); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		ret := []Group{}

		for _, value := range group {
			var temp Group
			temp.ID = value.ID
			temp.CompanionName = value.CompanionName
			temp.GroupName = value.GroupName
			ret = append(ret, temp)
		}
		c.JSON(http.StatusOK, gin.H {
			"success": true,
			"message": "query completed.",
			"data": ret,
		})
	})

	// untuk menambah grup baru di admin page
	r.POST("/group", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"success": false,
				"error": "get form err: " + err.Error(),
			})
			return
		}

		rand.Seed(time.Now().Unix())

		str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		shuff := []rune(str)

		rand.Shuffle(len(shuff), func(i, j int) {
			shuff[i], shuff[j] = shuff[j], shuff[i]
		}) 
		file.Filename = string(shuff); 

		if err := c.SaveUploadedFile(file, "./assets/"+file.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"Success": false, 
				"error": "upload file err: " + err.Error(),
			})
			return
		}

		godotenv.Load("../.env")
		newGroup := Model.Group {
			GroupName: c.PostForm("group_name"),
			LineGroup: c.PostForm("line_group"),
			CompanionName: c.PostForm("companion_name"),
			IDLine: c.PostForm("id_line"),
			LinkFoto: os.Getenv("BASE_URL") + "/assets/"+file.Filename,
		}

		if err := db.Create(&newGroup); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"success": false,
				"message": "error when inserting a new group",
				"error": err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H {
			"success": true,
			"message": "a new group has successfully created",
			"error": nil,
			"data": gin.H {
				"nama_grup": newGroup.GroupName,
				"nama_pendamping": newGroup.CompanionName,
			},
		})

	})

	// untuk mendapatkan data grup berdasarkan id
	r.GET("/group/:id", func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H {
				"Success": false,
				"message": "group_id is not available",
			})
			return
		}

		group := Model.Group{}

		if result := db.Where("id = ?", id).Preload("Student").Take(&group); result.Error != nil {
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
			"data": group,
		})
	})

	//untuk delete group berdasarkan id
	r.DELETE("/group/:id", func(c *gin.Context) {
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
		group := Model.Group{
			ID: uint(parsedId),
		}
		if result := db.Delete(&group); result.Error != nil {
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

	//untuk memperbarui data grup berdasarkan id yang dimiliki
	r.PATCH("/group/:id", func(c *gin.Context) { 
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H {
				"Success": false,
				"message": "id is not available",
			})
			return
		}

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"success": false,
				"error": "get form err: " + err.Error(),
			})
			return
		}

		rand.Seed(time.Now().Unix())

		str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		shuff := []rune(str)

		rand.Shuffle(len(shuff), func(i, j int) {
			shuff[i], shuff[j] = shuff[j], shuff[i]
		}) 
		file.Filename = string(shuff); 

		if err := c.SaveUploadedFile(file, "./assets/"+file.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"Success": false, 
				"error": "upload file err: " + err.Error(),
			})
			return
		}
		
		godotenv.Load("../.env")
		newGroup := Model.Group {
			GroupName: c.PostForm("group_name"),
			LineGroup: c.PostForm("line_group"),
			CompanionName: c.PostForm("companion_name"),
			IDLine: c.PostForm("id_line"),
			LinkFoto: os.Getenv("BASE_URL") + "/assets/"+file.Filename,
		}

		result := db.Where("id = ?", id).Model(&newGroup).Updates(newGroup)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H {
				"success": false,
				"message": "Error when updating the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		if result = db.Where("id = ?", id).Preload("Student").Take(&newGroup); result.Error != nil {
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
			"data":    newGroup,
		})
	})

	//untuk memperbarui nilai mahasiswa
	r.PATCH("/marking/:id", func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")

		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H {
				"success": false,
				"message": "id is not supplied",
			})
			return
		}

		body := Model.StudentMarking{}

		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"success": false,
				"message": "body is invalid",
				"error": err.Error(),
			})
			return
		}

		marking := Model.StudentMarking{
			MarkingID: body.MarkingID,
			StudentID: body.StudentID,
			Mark: body.Mark,
		}

		result := db.Where("id = ?", id).Model(&marking).Updates(&marking)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when updating marking data.",
				"error":   result.Error.Error(),
			})
			return
		}

		if result = db.Where("id = ?", id).Take(&marking); result.Error != nil {
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
				"message": "marking not found.",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "nilai berhasil diperbarui",
			"data":    marking,
		})
	})
}
