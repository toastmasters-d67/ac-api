package controllers

import (
	"api/models"
	"api/pkg"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

func NewTransactions() *Db {
	db := pkg.InitDb()
	db.AutoMigrate(&models.Transaction{})
	return &Db{Db: db}
}

func (db *Db) CreateTransaction(c *gin.Context) {
	c.MultipartForm()
	info := c.Request.PostForm["TradeInfo"][0]
	sha := c.Request.PostForm["TradeSha"][0]
	encoded := pkg.GetTradeSha(info)
	valid := sha == encoded
	decoded := pkg.DecodeTradeInfo(info)

	sj, err := simplejson.NewJson([]byte(decoded))
	if err != nil {
		fmt.Printf("err = %v \n", err)
	}
	message := sj.Get("Message").MustString()
	result := sj.Get("Result")

	item := models.Transaction{}
	item.ID = result.Get("MerchantOrderNo").MustString()
	item.Valid = valid
	item.Status = sj.Get("Status").MustString()
	item.Amount = result.Get("Amt").MustInt()
	item.Message = pkg.DecodeUnicode(message)
	item.TradeNo = result.Get("TradeNo").MustString()
	item.Ip = result.Get("IP").MustString()
	item.Bank = result.Get("EscrowBank").MustString()
	item.BankCode = result.Get("PayBankCode").MustString()
	item.Time = result.Get("PayTime").MustString()
	item.AccountCode = result.Get("PayerAccount5Code").MustString()
	item.CodeNo = result.Get("CodeNo").MustString()
	item.StoreType = result.Get("StoreType").MustInt()
	item.Store = result.Get("Store").MustString()
	fmt.Printf("item = %v \n", item)

	err = models.CreateTransaction(db.Db, &item)
	if err != nil {
		fmt.Printf("err = %v \n", err)
		return
	}
}

func (db *Db) GetTransactions(c *gin.Context) {
	transactions := []models.Transaction{}
	err := models.GetTransactions(db.Db, &transactions)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (db *Db) GetTransaction(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction := models.Transaction{}
	err := models.GetTransaction(db.Db, &transaction, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, transaction)
}
