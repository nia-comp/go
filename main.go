package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type REGISTER struct {
	USERNAME string `json:"username" binding:"required"`
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
	PHOTO    string `json:"photo" binding:"required"`
}

type User struct {
	ID       int
	USERNAME string
	EMAIL    string
	PASSWORD string
	PHOTO    string
	UpdatedAt string
	CreatedAt string
}

type SignIn struct {
	email string
	password string
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users/register", func(c *gin.Context) {
		var register REGISTER
		c.BindJSON(&register)

		hashed, _ := HashPassword(register.PASSWORD)

		user := User{USERNAME: register.USERNAME, EMAIL: register.EMAIL, PASSWORD: hashed, PHOTO: register.PHOTO}
		if err != nil {
        panic("failed to connect database")
    }
		db.Create(&user)
		c.JSON(200, gin.H{"nama": "hehe"})
	})

	r.POST("/users/signin", func (c *gin.Context)  {
		var signIn SignIn

		var result User
		db.Model(User{EMAIL: signIn.email}).First(&result)
		if result == (User{}) {
			c.JSON(404, gin.H{"error": "email tidak ketemu"})
		}

		password := result.PASSWORD
		if CheckPasswordHash(signIn.password, password) == true {
			c.JSON(200, gin.H{"status": "berhasil"})
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
