package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			panic("Cannot load env: " + err.Error())
		}
	}

	// connect to database
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		"10.101.114.115", // os.Getenv("MYSQL_HOST")
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	fmt.Println(22233)
	fmt.Println("111 dsn ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(3344)

	// setup server
	r := gin.Default()
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = []string{"*"}
	ginConfig.AllowCredentials = true
	ginConfig.AllowHeaders = []string{
		"Access-Control-Allow-Origin",
		"Origin",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"Referer",
		"X-Size",
	}
	r.Use(cors.New(ginConfig))

	// api
	group := r.Group("api/todos")

	group.GET("", func(c *gin.Context) {
		tasks := make([]Task, 0)
		_ = db.Model(&Task{}).Find(&tasks)

		c.JSON(200, tasks)
	})

	group.POST("", func(c *gin.Context) {
		var task Task
		c.Bind(&task)

		task.CreatedAt = time.Now()

		_ = db.Create(&task)

		c.JSON(200, task)
	})

	r.Run()
}

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}
