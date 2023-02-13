package controllers

import (
	"api/models"
	"api/pkg"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func NewCarts() *Db {
	db := pkg.InitDb()
	db.AutoMigrate(&models.Cart{})
	return &Db{Db: db}
}

func (db *Db) CreateCart(c *gin.Context) {
	merchantID := os.Getenv("MERCHANT_ID")
	version := os.Getenv("VERSION")
	seconds := time.Now().UTC().UnixNano() / 1e9
	timeStamp := strconv.FormatInt(seconds, 10)

	json := models.Item{}
	c.BindJSON(&json)
	params := map[string]string{
		"MerchantID":      merchantID,
		"Version":         version,
		"MerchantOrderNo": timeStamp,
		"TimeStamp":       timeStamp,
		"Amt":             strconv.Itoa(json.Amount),
		"RespondType":     "JSON",
		"ItemDesc":        json.Description,
		"Email":           json.Email,
		"ReturnURL":       json.Url + "/callback",
		"NotifyURL":       json.Url + "/notify",
	}
	content := pkg.GetTradeInfo(params)
	sha := pkg.GetTradeSha(content)

	cart := models.Cart{}
	cart.Amount = json.Amount
	cart.Email = json.Email
	cart.Description = json.Description
	fmt.Printf("cart: %v\n\n", cart)

	err := models.CreateCart(db.Db, &cart)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// c.JSON(http.StatusOK, cart)
	c.JSON(http.StatusOK, gin.H{
		"merchantID": merchantID,
		"version":    version,
		"content":    content,
		"sha":        sha,
	})
}

func (db *Db) GetCarts(c *gin.Context) {
	cart := []models.Cart{}
	err := models.GetCarts(db.Db, &cart)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, cart)
}

func (db *Db) GetCart(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cart := models.Cart{}
	err := models.GetCart(db.Db, &cart, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, cart)
}
