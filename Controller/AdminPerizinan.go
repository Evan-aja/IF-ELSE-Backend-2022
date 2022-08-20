package Controller

import (
	"ifelse/Auth"
	"ifelse/Model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminPerizinan(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")

	type Student struct {
		ID        uint   `json:"id"`
		Name      string `json:"name"`
		GroupName string `json:"group_name"`
		NIM       string `json:"nim"`
		LinkSurat string `json:"link_surat"`
	}

	r.GET("/admin/perizinan/:id", Auth.Authorization(), func(c *gin.Context) {
		// buat ambil ID untuk agenda
		id, _ := c.Params.Get("id")
		
		// buat ngambil ID dari authentication
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
			c.JSON(http.StatusForbidden, gin.H {
				"success": false,
				"message": "unauthorized access :(",
				"error": nil,
			})
			return
		}

		var perizinan []Model.Perizinan

		if result := db.Where("agenda_id = ?", id).Preload("Student").Preload("Agenda").Find(&perizinan); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}

		var group Model.Group
		var ret []Student

		for i := 0; i < len(perizinan); i++ {
			if result := db.Where("id = ?", perizinan[i].Student.GroupID).Find(&group); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Error when querying the database.",
					"error":   result.Error.Error(),
				})
				return
			}
			var temp Student
			temp.ID = perizinan[i].Student.ID
			temp.Name = perizinan[i].Student.Name
			temp.GroupName = group.GroupName
			temp.NIM = perizinan[i].Student.NIM
			temp.LinkSurat = perizinan[i].LinkSurat
			ret = append(ret, temp)
		}

		c.JSON(http.StatusOK, gin.H {
			"success": true,
			"message": "query completed.",
			"data":    ret,
		})
	})
}
