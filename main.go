package main

import (
	"api/controllers"
	"api/helpers"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Origin"}
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	Users := controllers.NewUsers()
	Orders := controllers.NewOrders()
	Transactions := controllers.NewTransactions()
	r.POST("/api/v1/user", Users.CreateUser)
	r.POST("/api/v1/login", Users.AuthenticateUser)
	r.POST("/api/v1/notify", Transactions.CreateTransaction)
	r.POST("/api/v1/callback", controllers.GetCallback)

	r.Use(helpers.VerifyToken)
	r.GET("/api/v1/user", Users.GetUser)
	r.GET("/api/v1/order", Users.GetOrders)
	r.POST("/api/v1/order", Orders.CreateOrder)
	return r
}

func main() {
	r := SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
