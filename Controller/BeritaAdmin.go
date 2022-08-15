package Controller

import (
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// hubungi mas fajri
func BeritaAdmin(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/admin/berita")

	// Get All Berita
	r.GET("/", func(c *gin.Context) {
		var berita []Model.Berita

		if err := db.Find(&berita).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "Cannot find data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"berita": berita,
		})
	})

	r.Static("/assets", "./assets")

	// pake :id
	r.PATCH("/baru/gambar/:id", func(c *gin.Context) {
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "user_id is not available",
			})
			return
		}

		file, err := c.FormFile("gambar")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "get form err: " + err.Error(),
			})
			return
		}

		if err := c.SaveUploadedFile(file, "./assets/"+file.Filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"error":   "upload file err: " + err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "File " + "./assets/" + file.Filename + " uploaded successfully",
		})

		inputGambar := Model.Berita{
			Gambar: file.Filename,
		}

		if err := db.Where("id = ?", id).Model(&inputGambar).Updates(inputGambar).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "cannot update data",
			})
			return
		}
	})

	// Buat Berita Baru
	r.POST("/baru", func(c *gin.Context) {
		var input Model.Berita

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "cannot get input",
			})
			return
		}

		if input.Judul == "" || input.Konten == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "judul dan konten tidak boleh kosong",
				"success": false,
			})
			return
		}

		input.IsPublished = true

		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "cannot create data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":      "Berita Berhasil Ditambahkan",
			"success":      true,
			"Title":        input.Judul,
			"Gambar":       input.Gambar,
			"Konten":       input.Konten,
			"Is Published": input.IsPublished,
		})
	})

	// Edit / Update berita
	r.PATCH("/edit/:id", func(c *gin.Context) {
		var input Model.Berita
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "url id is not available",
			})
			return
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "cannot get input",
			})
			return
		}

		if err := db.Where("id = ?", id).Model(&input).Updates(input).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "cannot update data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Berita Berhasil Diupdate",
			"success": true,
			"Title":   input.Judul,
			"Gambar":  input.Gambar,
			"Konten":  input.Konten,
		})
	})

	// Haput Berita
	r.DELETE("/hapus/:id", func(c *gin.Context) {
		var input Model.Berita
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "id is not available",
			})
			return
		}

		if err := db.Where("id = ?", id).Delete(&input).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"msg":     "cannot delete data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Berita Berhasil Dihapus",
			"success": true,
		})
	})
}
