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

func AdminGroup(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	// Group Routers
	type Group struct {
		ID            uint   `json:"id"`
		CompanionName string `json:"companion_name"`
		GroupName     string `json:"group_name"`
	}

	// untuk menampilkan seluruh grup yang ada
	r.GET("/admin/group", Auth.Authorization(), func(c *gin.Context) {
		var group []Model.Group

		if result := db.Where("group_name != ?", "Belum diatur").Preload("Student").Find(&group); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		var ret []Group

		for _, value := range group {
			var temp Group
			temp.ID = value.ID
			temp.CompanionName = value.CompanionName
			temp.GroupName = value.GroupName
			ret = append(ret, temp)
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    ret,
		})
	})
	r.Static("/admin/image", "./Images")

	// untuk menambah grup baru di admin page
	r.POST("/admin/group", Auth.Authorization(), func(c *gin.Context) {
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
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		} 	

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "get form err: " + err.Error(),
			})
			return
		}

		rand.Seed(time.Now().Unix())

		str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		shuff := []rune(str)

		rand.Shuffle(len(shuff), func(i, j int) {
			shuff[i], shuff[j] = shuff[j], shuff[i]
		})
		file.Filename = string(shuff)

		if err := c.SaveUploadedFile(file, "./Images/"+file.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"error":   "upload file err: " + err.Error(),
			})
			return
		}

		godotenv.Load("../.env")
		newGroup := Model.Group{
			GroupName:     c.PostForm("group_name"),
			LineGroup:     c.PostForm("line_group"),
			CompanionName: c.PostForm("companion_name"),
			IDLine:        c.PostForm("id_line"),
			LinkFoto:      os.Getenv("BASE_URL") + "/api/admin/image/" + file.Filename,
		}

		if err := db.Create(&newGroup); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when inserting a new group",
				"error":   err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "a new group has successfully created",
			"error":   nil,
			"data": gin.H{
				"nama_grup":       newGroup.GroupName,
				"nama_pendamping": newGroup.CompanionName,
			},
		})

	})

	// untuk mendapatkan data grup berdasarkan id
	r.GET("/admin/group/:id", Auth.Authorization(), func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "group is not available",
			})
			return
		}

		group := Model.Group{}

		if result := db.Where("id = ?", id).Take(&group); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		var student []Model.Student
		if result := db.Where("group_id = ?", id).Find(&student); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		for i := 0; i < len(student); i++ {
			if result := db.Where("id = ?", student[i].GroupID).Find(&group); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Error when querying the database.",
					"error":   result.Error.Error(),
				})
				return
			}
			student[i].GroupName = group.GroupName
		}

		group.Student = student


		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    group,
		})
	})

	//untuk memperbarui data grup berdasarkan id yang dimiliki
	r.PATCH("/admin/group/:id", Auth.Authorization(), func(c *gin.Context) {
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

		
		if user.RoleId != 3 && user.RoleId != 0 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
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

		file, _ := c.FormFile("file")

		var newGroup Model.Group
		var group Model.Group

		if file != nil {
			rand.Seed(time.Now().Unix())

			str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	
			shuff := []rune(str)
	
			rand.Shuffle(len(shuff), func(i, j int) {
				shuff[i], shuff[j] = shuff[j], shuff[i]
			})
			file.Filename = string(shuff)
	
			if err := c.SaveUploadedFile(file, "./Images/"+file.Filename); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Success": false,
					"error":   "upload file err: " + err.Error(),
				})
				return
			}
	
			godotenv.Load("../.env")
			newGroup = Model.Group{
				GroupName:     c.PostForm("group_name"),
				LineGroup:     c.PostForm("line_group"),
				CompanionName: c.PostForm("companion_name"),
				IDLine:        c.PostForm("id_line"),
				LinkFoto:      os.Getenv("BASE_URL") + "/api/admin/image/" + file.Filename,
			}
		} else {
			if res := db.Where("id = ?", id).Take(&group); res.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Error when updating the database.",
					"error":   res.Error.Error(),
				})
			}
			newGroup = Model.Group{
				GroupName:     c.PostForm("group_name"),
				LineGroup:     c.PostForm("line_group"),
				CompanionName: c.PostForm("companion_name"),
				IDLine:        c.PostForm("id_line"),
				LinkFoto:      group.LinkFoto,
			}
		}

		result := db.Where("id = ?", id).Model(&newGroup).Updates(newGroup)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
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

	//untuk delete group berdasarkan id
	r.DELETE("/admin/group/:id", Auth.Authorization(), func(c *gin.Context) {
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

		if user.RoleId != 3 && user.RoleId != 0 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
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

}
