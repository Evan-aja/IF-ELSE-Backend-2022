package Controller

import (
	"crypto/sha512"
	"encoding/hex"
	"ifelse/Model"
	"net/http"
	"os"
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
		regist := Model.Student{
			NIM:     input.NIM,
			Address: input.Address,
		}
		if err := db.Create(&regist); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong with student creation",
				"error":   err.Error.Error(),
			})
			return
		}
		regist2 := Model.User{
			Name:      input.Name,
			Username:  input.Username,
			Email:     input.Email,
			Password:  hash(input.Password),
			StudentID: regist.ID,
		}
		if err := db.Create(&regist2); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Something went wrong with user creation",
				"error":   err.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
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
		if err := db.Where("email=?", input.Email).Take(&email); err.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Email does not exist",
				"error":   err.Error.Error(),
			})
			return
		}
		if email.Password == hash(input.Password) {
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
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Welcome, here's your token. don't lose it ;)",
				"data": gin.H{
					"email": email.Email,
					"name":  email.Name,
					"token": strToken,
				},
			})
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Did you forget your own password?",
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
