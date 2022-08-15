package Controller

import (
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Pendataan(db *gorm.DB, q *gin.Engine) {

	r := q.Group("/api")

// Get All
	r.GET("/pendataan", func(c *gin.Context) {
		var pendataan []Model.Agenda

        if result := db.Find(&pendataan); result.Error != nil {
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
            "data": pendataan,
        })
	})
// Get By ID from Agenda
	r.GET("/pendataan/:id", func(c *gin.Context){
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "user_id is not available",
			})
			return
		}
		pendataan :=Model.Agenda{}
		if result := db.Where("id = ?", id).Take(&pendataan); result.Error != nil {
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
			"data":    pendataan,
		})


	})
// Get By ID from pendataan
	r.GET("/pendataann/:id", func(c *gin.Context){
		id, isIdExists := c.Params.Get("id")
		if !isIdExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"Success": false,
				"message": "user_id is not available",
			})
			return
		}
		pendataan :=Model.Pendataan{}
		if result := db.Where("id = ?", id).Take(&pendataan); result.Error != nil {
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
			"data":    pendataan,
		})



	})

	// Update
	r.PATCH("/pendataan/:id", func(c *gin.Context) {
        var input Model.Pendataan
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
            "Data":   input,
        })
    })



}
