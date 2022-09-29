package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type REGISTER struct {
	ID       int    `json:"id" binding:"required"`
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
	UpdatedAt time.Time
	CreatedAt time.Time
}

db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users/register", func(c *gin.Context) {
		var register REGISTER
		c.BindJSON(&register)
		c.JSON(200, gin.H{"nama": register.USERNAME})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func insert() {
	db.Ommit()
}
