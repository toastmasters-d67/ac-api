package controllers

import (
	"api/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Db models.Database

func GetCallback(c *gin.Context) {
	fmt.Println("callback")
	seconds := time.Now().UTC().UnixNano() / 1e9
	timeStamp := strconv.FormatInt(seconds, 10)
	fmt.Printf("timeStamp: %v\n\n", timeStamp)
	fmt.Printf("time.Now(): %v\n\n", time.Now())
	c.MultipartForm()
	for key, value := range c.Request.PostForm {
		fmt.Printf("%v = %v \n", key, value)
	}
}
