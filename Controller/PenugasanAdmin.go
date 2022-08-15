package Controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"ifelse/Model"
	"net/http"
)

// hubungi mas erycson
func PenugasanAdmin(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/admin/penugasan")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Penugasan Admin If-Else 2022",
			"success": true,
		})
	})

	// Get All Penugasan
	r.GET("/", func(c *gin.Context) {
		var penugasan []Model.Penugasan

		if err := db.Find(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}

		if penugasan == nil {
			c.JSON(http.StatusOK, gin.H{
				"data":    penugasan,
				"Error":   "Data Is Empty",
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    penugasan,
			"success": true,
		})
	})

	// Get Penugasan By ID
	r.GET("/:id", func(c *gin.Context) {
		var penugasan Model.Penugasan

		id := c.Param("id")

		if err := db.Where("id = ?", id).First(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}

		if penugasan.Title == "" {
			c.JSON(http.StatusOK, gin.H{
				"data":    penugasan,
				"Error":   "Data Is Empty",
				"success": false,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    penugasan,
			"success": true,
		})
	})

	// Create Penugasan
	r.POST("/baru", func(c *gin.Context) {
		var penugasan Model.Penugasan

		//inputLink := [2]string{penugasan.Link1, penugasan.Link2}
		//var jsonData = [2]byte(inputLink)

		if err := c.BindJSON(&penugasan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't read input",
			})
			return
		}

		if penugasan.Title == "" || penugasan.Description == "" || penugasan.Condition == "" || penugasan.Step == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "All Field Must Be Filled",
				"success": false,
			})
			return
		}

		//inputLink := [2]string{penugasan.Link1, penugasan.Link2}

		//penugasan.Link[0] = penugasan.Link1
		//penugasan.Link[1] = penugasan.Link2

		if err := db.Create(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't create data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    penugasan,
			"success": true,
		})
	})

	// Update Penugasan
	r.PATCH("/edit/:id", func(c *gin.Context) {
		var penugasan Model.Penugasan

		id := c.Param("id")

		if err := db.Where("id = ?", id).First(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't create data",
			})
			return
		}

		if err := c.BindJSON(&penugasan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't read input",
			})
			return
		}

		if err := db.Save(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't update data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    penugasan,
			"success": true,
		})
	})

	// Delete Penugasan by Id
	r.DELETE("/delete", func(c *gin.Context) {
		var penugasan Model.Penugasan

		if err := c.BindJSON(&penugasan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't read input",
			})
			return
		}

		id := penugasan.ID

		if err := db.Where("id = ?", id).First(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't find data",
			})
			return
		}

		if err := db.Delete(&penugasan).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"success": false,
				"Reason":  "Can't delete data",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data":    penugasan,
			"success": true,
		})
	})

	// Delete All Penugasan
	//r.DELETE("/delete/all", func(c *gin.Context) {
	//	var penugasan []Model.Penugasan
	//
	//	if err := db.Find(&penugasan).Error; err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": err.Error(),
	//			"success": false,
	//			"Reason":  "Can't find data",
	//		})
	//		return
	//	}
	//
	//	if err := db.Delete(&penugasan).Error; err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": err.Error(),
	//			"success": false,
	//			"Reason":  "Can't delete data",
	//		})
	//		return
	//	}

	//	c.JSON(http.StatusOK, gin.H{
	//		"data":    penugasan,
	//		"success": true,
	//	})
	//})

	// Delete All Penugasan
	//r.DELETE("/delete/all/:id", func(c *gin.Context) {
	//	var penugasan []Model.Penugasan
	//
	//	id := c.Param("id")
	//
	//	if err := db.Where("id = ?", id).Find(&penugasan).Error; err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": err.Error(),
	//			"success": false,
	//			"Reason":  "Can't find data",
	//		})
	//		return
	//	}
	//
	//	if err := db.Delete(&penugasan).Error; err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"message": err.Error(),
	//			"success": false,
	//			"Reason":  "Can't delete data",
	//		})
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"data":    penugasan,
	//		"success": true,
	//	})
	//})
}
