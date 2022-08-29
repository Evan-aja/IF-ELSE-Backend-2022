package Controller

import (
	"crypto/sha512"
	"encoding/hex"
	"ifelse/Model"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api")
	// /api/register user baru
	r.POST("/register", func(c *gin.Context) {
		var input Model.Register
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error(),
			})
			return
		}
		var regexNIM = regexp.MustCompile(`([0-9])\w+515020+([0-9])\w+`)
		var regexNIMKhusus = regexp.MustCompile(`([0-9])\w+515021+([0-9])\w+`)
		var isMatchNIM = regexNIM.MatchString(input.NIM)
		var isMatchNIMKhusus = regexNIMKhusus.MatchString(input.NIM)
		var regexEmail = regexp.MustCompile(`([\p{L}\d])\w+@student\.ub\.ac\.id`)
		var isMatchEmail = regexEmail.MatchString(input.Email)
		var regist Model.Student
		if isMatchNIM || isMatchNIMKhusus {
			regist = Model.Student{
				Name: input.Name,
				NIM:  input.NIM,
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H {
				"success":false,
				"message": "format NIM salah",
				"error": nil,
			})
			return
		}
		if err := db.Create(&regist); err.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "duplicate entry for NIM!",
				"error":   err.Error.Error(),
			})
			return
		}

		var regist2 Model.User
		if isMatchEmail {
			regist2 = Model.User{
				Username:  input.Username,
				Email:     input.Email,
				Password:  hash(input.Password),
				StudentID: regist.ID,
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H {
				"success":false,
				"message": "format email Anda salah",
				"error": nil,
			})
			db.Where("id = ?", regist.ID).Delete(regist)
			return
		}

		if err := db.Create(&regist2); err.Error != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "duplicate entry for username!",
				"error":   err.Error.Error(),
			})
			db.Where("id = ?", regist.ID).Delete(regist)
			return
		}

		var allAgenda []Model.Agenda
		if result := db.Find(&allAgenda); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when querying the database.",
				"error":   result.Error.Error(),
			})
			return
		}
		
		// check agenda is nil or not
		if allAgenda != nil {
			// assign marking to `mahasiswa`
			mark := make([]Model.Marking, len(allAgenda))

			for i, element := range allAgenda {
				mark[i].AgendaID = element.ID
				mark[i].StudentID = regist2.StudentID
				mark[i].ID = 0
				if err := db.Create(&mark[i]).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "can't create marks",
						"success": false,
						"error":   err.Error(),
					})
					return
				}
			}	
		}

		// var allTask []Model.Task
		
		// if result := db.Find(&allTask); result.Error != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"success": false,
		// 		"message": "Error when querying the database.",
		// 		"error":   result.Error.Error(),
		// 	})
		// 	return
		// }

		// if allTask != nil {
		// 	var allLink []Model.Links

		// 	studentTask := make([]Model.StudentTask, len(allTask))
			
		// 	var linkId []uint
		// 	// assign link ke siswa
		// 	for i := 0; i < len(allTask); i++ {
		// 		studentTask[i].StudentID = regist2.StudentID
		// 		studentTask[i].TaskID = allTask[i].ID
		// 		if result := db.Where("task_id = ?", studentTask[i].TaskID).Find(&allLink); result.Error != nil {
		// 			c.JSON(http.StatusInternalServerError, gin.H{
		// 				"success": false,
		// 				"message": "Error when querying the database.",
		// 				"error":   result.Error.Error(),
		// 			})
		// 			return
		// 		}
		// 		linkId = append(linkId, allLink[i].ID)
		// 		fmt.Println(linkId)
				
		// 		for j := 0; j < int(allTask[i].JumlahLink); j++ {
		// 			studentTask[i].LinkPos = int32(j)
		// 			studentTask[i].LinkID = linkId[j]
		// 			studentTask[i].ID = 0
		// 			if err := db.Create(&studentTask[i]).Error; err != nil {
		// 				c.JSON(http.StatusInternalServerError, gin.H{
		// 					"message": "can't create links",
		// 					"success": false,
		// 					"error":   err.Error(),
		// 				})
		// 				return
		// 			}
		// 		}
		// 	}
		// }
		
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Account created successfully",
			"error":   nil,
			"data": gin.H{
				"nama":  input.Name,
				"email": input.Email,
			},
		})
	})
	// /api/login user terdaftar
	r.POST("/login", func(c *gin.Context) {
		var input Model.Register
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Something went wrong",
				"error":   err.Error(),
			})
			return
		}
		email := Model.User{}
		if input.Username == "" {
			if err := db.Where("email=?", input.Email).Take(&email); err.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "email atau password Anda salah.",
					"error":   err.Error.Error(),
				})
				return
			}
		} else if input.Email == "" {
			if err := db.Where("username=?", input.Username).Take(&email); err.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Email does not exist",
					"error":   err.Error.Error(),
				})
				return
			}
		}

		if email.Password == hash(input.Password)  {
			token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"id":  email.ID,
				"exp": time.Now().Add(time.Hour * 7 * 24).Unix(),
			})
			godotenv.Load("../.env")
			strToken, err := token.SignedString([]byte(os.Getenv("TOKEN_G")))
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "Something went wrong",
					"error":   err.Error(),
				})
				return
			}
			var IsAdmin bool
			if email.RoleId > 3 {
				IsAdmin = false
			} else {
				IsAdmin = true
			}
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Welcome, here's your token. don't lose it ;)",
				"data": gin.H{
					"username":email.Username,
					"IsAdmin": IsAdmin,
					"email": email.Email,
					"token": strToken,
				},
			})
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "email atau password Anda salah.",
			})
			return
		}
	})

}

// hash sha512 untuk password
func hash(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))
	pass := hex.EncodeToString(hash.Sum(nil))
	return pass
}
