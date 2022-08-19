package Controller

import (
	"fmt"
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

func AdminAgenda(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	r.Static("/agenda/image", "./Images")
	type Agenda struct {
		ID      uint   `json:"id"`
		Title   string `json:"title"`
		StartAt string `json:"start_at"`
		EndAt   string `json:"end_at"`
	}

	// post a new agenda
	r.POST("/admin/agenda", Auth.Authorization(), func(c *gin.Context) {
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

		if user.RoleId != 0 && user.RoleId != 2 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}
		image, err := c.FormFile("image")
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
		image.Filename = string(shuff)

		godotenv.Load("../.env")

		if err := c.SaveUploadedFile(image, "./Images/"+image.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"error":   "upload file err: " + err.Error(),
			})
			return
		}

		newAgenda := Model.Agenda{
			Title:            c.PostForm("title"),
			Content:          c.PostForm("content"),
			Image:            os.Getenv("BASE_URL") + "/api/agenda/image/" + image.Filename,
			StartAt:          c.PostForm("start_at"),
			EndAt:            c.PostForm("end_at"),
			PerizinanStartAt: c.PostForm("perizinan_start_at"),
			PerizinanEndAt:   c.PostForm("perizinan_end_at"),
		}

		if err := db.Create(&newAgenda); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when inserting a new agenda",
				"error":   err.Error.Error(),
			})
			return
		}

		var student []Model.Student
		if err := db.Find(&student); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error when querying database",
				"success": false,
				"error":   err.Error.Error(),
			})
		}

		var agenda []Model.Agenda
		if err := db.Find(&agenda); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error when querying database",
				"success": false,
				"error":   err.Error.Error(),
			})
		}

		var mark Model.Marking

		mark.AgendaID = newAgenda.ID
		for i := 0; i < len(student); i++ {
			fmt.Println(len(student))
			mark.StudentID = student[i].ID
			mark.ID = 0
			if err := db.Create(&mark).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "can't create links",
					"success": false,
					"error":   err.Error(),
				})
				return
			}
		}

		c.JSON(http.StatusCreated, gin.H{
			"success":     true,
			"message":     "a new agenda has successfully created",
			"error":       nil,
			"nama_agenda": newAgenda.Title,
		})
	})

	// get all agenda
	r.GET("/admin/agenda", Auth.Authorization(),func(c *gin.Context) {
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

		if user.RoleId > 3 {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "unauthorized access :(",
				"error":   nil,
			})
			return
		}
		var agenda []Model.Agenda
		if result := db.Find(&agenda); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}
		var ret []Agenda

		for _, value := range agenda {
			var temp Agenda
			temp.ID = value.ID
			temp.Title = value.Title
			temp.StartAt = value.StartAt
			temp.EndAt = value.EndAt
			ret = append(ret, temp)
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    ret,
		})
	})

	// untuk mendapatkan detail agenda dari id
	r.GET("/admin/agenda/:id", Auth.Authorization(), func(c *gin.Context) {
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

		if user.RoleId > 3 {
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

		var agenda Model.Agenda

		if result := db.Where("id = ?", id).Take(&agenda); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query successful.",
			"error":   nil,
			"data":    agenda,
		})

	})

	// patch agenda by `agenda.id`
	r.PATCH("/admin/agenda/:id", Auth.Authorization(), func(c *gin.Context) {
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

		if user.RoleId != 0 && user.RoleId != 2 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		} 	
		id, _ := c.Params.Get("id")

		image, err := c.FormFile("image")
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
		image.Filename = string(shuff)

		godotenv.Load("../.env")

		if err := c.SaveUploadedFile(image, "./Images/"+image.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"error":   "upload file err: " + err.Error(),
			})
			return
		}

		var newAgenda Model.Agenda
		parsedId, _ := strconv.ParseUint(id, 10, 32)
		b1, _ := strconv.ParseBool(c.PostForm("is_published"))
		newAgenda = Model.Agenda{
			ID:               uint(parsedId),
			Title:            c.PostForm("title"),
			Content:          c.PostForm("content"),
			Image:            os.Getenv("BASE_URL") + "/api/agenda/image/" + image.Filename,
			StartAt:          c.PostForm("start_at"),
			EndAt:            c.PostForm("end_at"),
			PerizinanStartAt: c.PostForm("perizinan_start_at"),
			PerizinanEndAt:   c.PostForm("perizinan_end_at"),
			IsPublished:      b1,
		}

		if err := db.Where("id = ?", id).Model(&newAgenda).Select("*").Updates(newAgenda); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when inserting a new agenda",
				"error":   err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "a new agenda has successfully created",
			"error":   nil,
			"data":    newAgenda,
		})
	})

	// patch isPublished agenda
	r.PATCH("/admin/toggle-agenda/:id", Auth.Authorization(), func(c *gin.Context) {
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

		if user.RoleId != 0 && user.RoleId != 2 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		} 	
		id, _ := c.Params.Get("id")

		parsedId, _ := strconv.ParseUint(id, 10, 32)
		b1, _ := strconv.ParseBool(c.PostForm("is_published"))
		patchToggle := Model.Agenda{
			ID: uint(parsedId),
			IsPublished: b1,
		}

		if err := db.Where("id = ?", id).Model(&patchToggle).Select("is_published").Updates(patchToggle); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when inserting a new agenda",
				"error":   err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "a new agenda has successfully published",
			"error":   nil,
			"data":    patchToggle.IsPublished,
		})

		
	})
	
	// delete agenda by `agenda.id`
	r.DELETE("/admin/agenda/:id", Auth.Authorization(),func(c *gin.Context) {
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

		if user.RoleId != 0 && user.RoleId != 2 {
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
		parsedId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "ID is invalid.",
			})
			return
		}
		agenda := Model.Agenda{
			ID: uint(parsedId),
		}
		if result := db.Delete(&agenda); result.Error != nil {
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
