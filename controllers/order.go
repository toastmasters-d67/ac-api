package controllers

import (
	"api/helpers"
	"api/models"
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

func NewOrders() *Db {
	db := helpers.InitDb()
	db.AutoMigrate(&models.Order{})
	return &Db{Db: db}
}

func (db *Db) CreateOrder(c *gin.Context) {
	id, email, ok := helpers.GetSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	merchantID := os.Getenv("PAY_MERCHANT_ID")
	version := os.Getenv("PAY_VERSION")
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
		"Email":           email,
		"NotifyURL":       json.Url + "/notify",
		"ReturnURL":       json.Callback,
	}
	content := helpers.GetTradeInfo(params)
	sha := helpers.GetTradeSha(content)

	order := models.Order{}
	order.ID = timeStamp
	order.Amount = json.Amount
	order.Early = json.Early
	order.Double = json.Double
	order.First = json.First
	order.Second = json.Second
	order.Banquet = json.Banquet
	order.UserID = id
	order.Email = email
	order.Description = json.Description
	fmt.Printf("order: %v\n\n", order)

	err := models.CreateOrder(db.Db, &order)
	if err != nil {
		fmt.Printf("err: %v\n\n", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	url := os.Getenv("PAY_URL")
	c.JSON(http.StatusOK, gin.H{
		"merchantID": merchantID,
		"version":    version,
		"content":    content,
		"sha":        sha,
		"url":        url,
	})
}

func (db *Db) GetOrders(c *gin.Context) {
	_, email, ok := helpers.GetSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	orders, err := models.GetOrdersByEmail(db.Db, email)
	if err != nil {
		fmt.Printf("GetOrders() err: %v\n\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, orders)

}

func (db *Db) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := models.GetOrderById(db.Db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, order)
}
