package Controller

import (
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserNews(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	type News struct {
		ID      uint   `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// get all news (show title and content)
	r.GET("/news", func(c *gin.Context) {
		var news []Model.News

		if res := db.Find(&news); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "news is not found.",
				"error":   res.Error.Error(),
			})
			return
		}

		var ret []News

		for _, value := range news {
			var temp News
			temp.ID = value.ID
			temp.Title = value.Title
			temp.Content = value.Content
			ret = append(ret, temp)
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Search successful",
			"data":    ret,
		})
	})

	// get a specific news by id
	r.GET("/news/:id", func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		var news Model.News

		if result := db.Where("id = ?", id).Take(&news); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "students is not found.",
				"error":   result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Search successful",
			"data":    news,
		})
	})

	// get a latest news
	r.GET("/latest-news", func(c *gin.Context) {
		var news Model.News
		if result := db.Last(&news); result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "news is not found.",
				"error":   result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Search successful",
			"data":    news,
		})
	})

	// search news using `title` query
	r.POST("/news", func(c *gin.Context) {
		title, _ := c.GetQuery("q")

		var queryNews []Model.News

		if res := db.Where("title LIKE ?", "%"+title+"%").Find(&queryNews); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "news is not found.",
				"error":   res.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Search successful",
			"data":    queryNews,
		})

	})
}