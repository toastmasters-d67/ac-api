package main

import (
	"api/controllers"
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
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	Carts := controllers.NewCarts()
	Transactions := controllers.NewTransactions()
	r.POST("/api/v1/cart", Carts.CreateCart)
	r.POST("/api/v1/notify", Transactions.CreateTransaction)
	r.POST("/api/v1/callback", controllers.GetCallback)
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
