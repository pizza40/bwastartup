package main

import (
	"bwastartup/handler"
	"bwastartup/user"

	//"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//How It Works

//input dari user
//handler, mapping input dari user -> struct input
//service = melakukan mapping dari struct input ke struct user
//repository


func main(){
	//buat connection ke database mysql
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	
	

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	router.Run()
	

	//make router
	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()
}

//function handler (just like controller in laravel)
//handler gin
// func handler(c *gin.Context){
// 	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
//   	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User //insert data user to variable users
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users) //show data as JSON
// }