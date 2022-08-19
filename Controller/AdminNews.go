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

func AdminNews(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	r.Static("/admin/news", "./Images")
	type News struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		IsPublished bool   `json:"status"`
	}

	// untuk menambah berita baru di admin page
	r.POST("/admin/news", Auth.Authorization(), func(c *gin.Context) {
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

		x := user.RoleId != 0
		if x || user.RoleId != 2 {
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		} 
		file, err := c.FormFile("image")
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
		newNews := Model.News{
			Title:   c.PostForm("title"),
			Image:   os.Getenv("BASE_URL") + "/api/admin/news/" + file.Filename,
			Content: c.PostForm("content"),
		}

		if err := db.Create(&newNews); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "error when inserting a new news",
				"error":   err.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "a new group has successfully created",
			"error":   nil,
			"judul":   newNews.Title,
		})

	})

	// untuk mendapatkan seluruh data news
	r.GET("/admin/news", Auth.Authorization(), func(c *gin.Context) {
		var news []Model.News

		if result := db.Find(&news); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		var ret []News

		for _, value := range news {
			var temp News
			temp.ID = value.ID
			temp.Title = value.Title
			temp.IsPublished = value.IsPublished
			ret = append(ret, temp)
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "query completed.",
			"data":    ret,
		})
	})

	// untuk mendapatkan detail news dari id
	r.GET("/admin/news-detail/:id", Auth.Authorization(), func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "id is not available",
			})
			return
		}

		var news Model.News

		if result := db.Where("id = ?", id).Take(&news); result.Error != nil {
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
			"data":    news,
		})

	})

	// untuk memperbarui data news
	r.PATCH("/admin/news/:id", Auth.Authorization(), func(c *gin.Context) {
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

		x := user.RoleId != 0
		if x || user.RoleId != 2 {
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

		file, err := c.FormFile("image")
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

		if err := c.SaveUploadedFile(file, "./images/"+file.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"error":   "upload file err: " + err.Error(),
			})
			return
		}
		var newNews Model.News

		b1, _ := strconv.ParseBool(c.PostForm("is_published"))
		parsedId, _ := strconv.ParseUint(id, 10, 32)
		godotenv.Load("../.env")
		newNews = Model.News{
			ID:          uint(parsedId),
			Title:       c.PostForm("title"),
			Image:       os.Getenv("BASE_URL") + "/api/admin/news/" + file.Filename,
			Content:     c.PostForm("content"),
			IsPublished: b1,
		}

		result := db.Where("id = ?", id).Model(&newNews).Select("*").Updates(newNews)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when updating the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		if result = db.Where("id = ?", id).Take(&newNews); result.Error != nil {
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
			"data":    newNews,
		})
	})

	// untuk menghapus data news
	r.DELETE("/admin/news/:id", Auth.Authorization(), func(c *gin.Context) {
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

		x := user.RoleId != 0
		if x || user.RoleId != 2 {
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
		news := Model.News{
			ID: uint(parsedId),
		}
		if result := db.Delete(&news); result.Error != nil {
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
